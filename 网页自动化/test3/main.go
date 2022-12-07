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
	wd.Get("https://www.baidu.com/s?wd=go")
	time.Sleep(3 * time.Second)
	// 通过class属性定位元素
	ele1, _ := wd.FindElement(selenium.ByClassName, "s_ipt")
	// 清空网页元素的文本内容
	ele1.Clear()
	// 往网页元素输入文本内容
	ele1.SendKeys("Golang")
	time.Sleep(3 * time.Second)
	ele2, _ := wd.FindElement(selenium.ByID, "su")
	// 鼠标移动到网页元素
	ele2.MoveTo(1129, 15)
	time.Sleep(3 * time.Second)
	// 点击网页元素
	ele2.Click()
	time.Sleep(3 * time.Second)
	// 点击网页元素，Submit()用于表单按钮的点击
	ele2.Submit()
	// 获取网页元素的HTML标签
	tag, _ := ele2.TagName()
	fmt.Printf("获取网页元素的HTML标签：%v\n", tag)
	// 判断网页元素是否被选中
	// 通常用于checkbox和radio标签，返回值为true或false
	r, _ := ele2.IsSelected()
	fmt.Printf("判断网页元素是否被选中：%v\n", r)
	// 判断网页元素是否可编辑或可点击状态，返回值为true或false
	r1, _ := ele2.IsEnabled()
	fmt.Printf("判断网页元素是否可编辑或可点击状态：%v\n", r1)
	// 判断网页元素是否可见，返回值为true或false
	r2, _ := ele2.IsDisplayed()
	fmt.Printf("判断网页元素是否可见：%v\n", r2)
	ga, _ := ele2.GetAttribute("class")
	fmt.Printf("获取网页元素的属性class的值：%v\n", ga)
	time.Sleep(3 * time.Second)
	// 网页元素重新定位，因为上述点击操作使HTML代码发生变化
	ele2, _ = wd.FindElement(selenium.ByID, "su")
	// 获取网页元素的坐标位置
	p, _ := ele2.Location()
	fmt.Printf("获取网页元素的坐标位置：%v\n", p)
	// 网页元素显示在网页并获取坐标位置
	p1, _ := ele2.LocationInView()
	fmt.Printf("网页元素显示在网页并获取坐标位置：%v\n", p1)
	// 获取网页元素的大小
	s1, _ := ele2.Size()
	fmt.Printf("获取网页元素的大小：%v\n", s1)
	// 获取CSS样式的属性值，font-size是CSS样式名称
	c, _ := ele2.CSSProperty("font-size")
	fmt.Printf("获取CSS样式的属性值：%v\n", c)
	// 样式截图，返回值[]byte是图片的字节数据
	b, _ := ele2.Screenshot(true)
	// 保存图片
	ioutil.WriteFile("aa.jpg", b, 0755)
}
