package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
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
	url := "https://zhidao.baidu.com/question/496886588283871492.html"
	wd.Get(url)

	/* 切换iframe标签 */
	// iframe标签里面是嵌套HTML网页，不同HTML代码需要SwitchFrame切换
	// 通过iframe标签的id属性切换
	wd.SwitchFrame("MstReqFrm")
	// 定位iframe标签里面HTML网页的网页元素p
	e1, _ := wd.FindElement(selenium.ByTagName, "p")
	e1.SendKeys("b")
	// 参数frame设为nil则切换主HTML
	wd.SwitchFrame(nil)

	/* 切换浏览器多标签页 */
	// 获取浏览器所有标签页
	ss, _ := wd.WindowHandles()
	fmt.Printf("获取浏览器所有标签页：%v\n", ss)
	// 获取浏览器当前正在显示的标签页
	cs, _ := wd.CurrentWindowHandle()
	fmt.Printf("获取浏览器当前正在显示的标签页：%v\n", cs)
	// 切换标签页
	wd.SwitchWindow(ss[len(ss)-1])
	// 获取浏览器当前正在显示的标签页
	cs1, _ := wd.CurrentWindowHandle()
	fmt.Printf("获取浏览器当前正在显示的标签页：%v\n", cs1)
	time.Sleep(3 * time.Second)
}
