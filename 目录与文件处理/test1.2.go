package main

import (
	"fmt"
	"os"
	"time"
)

// 操作目录与文件
func main() {
	gw, _ := os.Getwd()
	fmt.Printf("获取当前目录：%v\n", gw)
	//改变当前工作目录
	os.Chdir("/Users/yanghaipeng/go/src/github.com/demo")
	gwn, _ := os.Getwd()
	fmt.Printf("改变当前工作目录：%v\n", gwn)
	// 创建文件
	f1, _ := os.Create("./1.txt")
	f1.Close()
	// 第二个参数mode在Windows系统下
	// mode为0200代表所有者可写
	// mode为0400代表只读
	// mode为0600代表读写
	//os.Chmod("D:/1.txt", 0400)
	// 修改文件的访问时间和修改时间
	nows := time.Now().Add(time.Hour)
	os.Chtimes("/Users/yanghaipeng/go/src/github.com/demo/1.txt", nows, nows)
	// 把字符串中带${var}或$var替换成指定指符串
	s := "你好，${1}${2}$3"
	fmt.Printf(os.Expand(s, func(k string) string {
		mapp := map[string]string{"1": "我是",
			"2": "go",
			"3": "语言"}
		return mapp[k]

	}))
	//创建目录
	os.Mkdir("/Users/yanghaipeng/go/src/github.com/demo/abc", os.ModePerm)
	//创建多级目录
	os.MkdirAll("/Users/yanghaipeng/go/src/github.com/demo/abc/d/e/f", os.ModePerm)
	// 删除文件或目录
	os.Remove("/Users/yanghaipeng/go/src/github.com/demo/abc/d/e/f")
	// 删除指定目录下所有文件
	os.RemoveAll("/Users/yanghaipeng/go/src/github.com/demo/abc")
	// 重命名文件
	os.Rename("/Users/yanghaipeng/go/src/github.com/demo/1.txt", "/Users/yanghaipeng/go/src/github.com/demo/1_new.txt")
	// 判断文件相同。Stat()获取文件信息，SameFile()判断文件相同
	f2, _ := os.Create("/Users/yanghaipeng/go/src/github.com/demo/2.txt")
	fs2, _ := f2.Stat()
	f3, _ := os.Create("/Users/yanghaipeng/go/src/github.com/demo/3.txt")
	fs3, _ := f3.Stat()
	fmt.Printf("f2和f3是否同一文件：%v\n", os.SameFile(fs2, fs3))
	// 返回临时目录
	fmt.Printf("返回临时目录：%v\n", os.TempDir())
}
