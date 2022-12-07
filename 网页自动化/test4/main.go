package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"io/ioutil"
	"time"
)

// 设置常量
const (
	// ChromeDriver路径，存放在E:\mygo，因此使用相对路径表示
	chromeDriver = "./Chromedriver/chromedriver"
	// ChromeDriver运行端口
	port = 8080
)

func main() {
	/* 开启WebDriver服务 */
	s, _ := selenium.NewChromeDriverService(chromeDriver, port)
	//关闭服务
	defer s.Stop()
	/* 连接WebDriver服务 */
	caps := selenium.Capabilities{}
	// 设置chrome特定功能
	chromeCaps := chrome.Capabilities{
		// 使用开发者调试模式
		ExcludeSwitches: []string{"enable-automation"},
	}
	// 将谷歌特定浏览器功能chromeCaps添加到caps
	caps.AddChrome(chromeCaps)
	// 根据浏览器功能连接Selenium
	urlPrefix := fmt.Sprintf("http://127.0.0.1:%d/wd/hub", port)
	wd, _ := selenium.NewRemote(caps, urlPrefix)
	// 关闭浏览器对象
	defer wd.Quit()
	// 访问网址
	wd.Get("https://www.baidu.com/s?wd=golang")
	time.Sleep(3 * time.Second)
	// 获取当前网址
	url, _ := wd.CurrentURL()
	fmt.Printf("当前URL地址：%v\n", url)
	//网页截图
	b, _ := wd.Screenshot()
	_ = ioutil.WriteFile("aa.jpg", b, 0755)
	// 获取网页标题
	t, _ := wd.Title()
	fmt.Printf("获取网页标题：%v\n", t)
	// 刷新网页
	wd.Refresh()
	// 执行js脚本实现网页元素操作
	e, _ := wd.FindElement(selenium.ByID, "kw")
	wd.ExecuteScript("arguments[0].click();", []interface{}{e})

}
