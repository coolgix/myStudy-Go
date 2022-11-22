package main

import (
	"fmt"
	"strings"
)

//func main() {
//	var a string
//	var b string
//	a = "hello world\n"
//	b = `hello
//		\n
//		world`
//
//	fmt.Printf("双引号的字符串：%v\n", a)
//	fmt.Printf("反引号的字符串：%v\n", b)
//
//}

//func main() {
//	fmt.Println("Tom\tjack")
//	fmt.Println("你\n好")
//	fmt.Println("E:\\go\\chapter3.6.go")
//	fmt.Println("Alic说：\"I love you\"")
//	fmt.Println("hello world")
//
//}

//func main() {
//
//	//输出字符a的ASCII
//	fmt.Printf("格式化符号v：%v\n", 'a')
//	//输出整形的数据类型
//	fmt.Printf("格式化符哈T：%T\n", 123)
//	//输出带双引号的字符串
//	fmt.Printf("格式化符号q：%q\n", "hello go")
//	//输出字符串的数据类型
//	fmt.Printf("格式化符号s: %s\n", "hello go")
//	//输出保留小数点的两位的浮点数，.2是小数点后保留的位数
//	fmt.Printf("格式化符号符f:%.2f\n", 1.223232131)
//	//输出十进制类型的整形
//	fmt.Printf("格式化符号d:%d\n", 31221)
//
//}

//func main() {
//	n := "hello world"
//	m := "i am Tom"
//	//使用"+"进行拼接
//	j := n + "," + m
//	fmt.Println(j)
//
//	//使用fmt.spreintf()拼接字符串
//	k := fmt.Sprintf("%s,%s", n, m)
//	fmt.Println(k)
//
//	//使用strings.Builder链接字符串
//	var builder strings.Builder
//	builder.WriteString(n)
//	builder.WriteString(",")
//	builder.WriteString(m)
//	fmt.Println(builder.String())
//
//	//使用bytes.Buffer链接字符串
//	var buffer bytes.Buffer
//	buffer.WriteString(n)
//	buffer.WriteString(",")
//	buffer.WriteString(m)
//	fmt.Println(buffer.String())
//
//}

//func main() {
//	n := "hello world"
//	m := "你好，go语言"
//	fmt.Println("字符串n的长度：", len(n))
//	fmt.Println("字符串m的长度：", len(m))
//}

//func main() {
//	n := "hello world"
//	m := "你好，go语言"
//
//	//使用utf8.RuneCountInstring()获取字符串的长度
//	fmt.Println("字符串n的长度：", utf8.RuneCountInString(n))
//	fmt.Println("字符串m的长度：", utf8.RuneCountInString(m))
//
//	//使用[]rune()获取字符串的长度
//	fmt.Println("字符串n的长度：", len([]rune(n)))
//	fmt.Println("字符串m的长度：", len([]rune(m)))
//}

//func main() {
//	n := "hi,go语言"
//	m := []rune(n)
//	for i := 0; i < len(n); i++ {
//		//%c输出每个字符
//		//%d输出字符对应的十进制值
//		fmt.Printf("%c-%d\n", m[i], m[i])
//	}
//}

//func main() {
//	n := "hi,go语言"
//	for _, s := range n {
//		//%c输出每个字符
//		//%d输出字符对应的十进制值
//		fmt.Printf("%c-%d\n", s, s)
//	}
//}

//func main() {
//	n := "hello-world-你好呀"
//	//获取字符串的子字符串world最开始的位置
//	m := strings.Index(n, "world") //m位置从左0到右第6个
//	fmt.Println("获取子字符串world的最开始的位置：", m)
//
//	//获取字符串的子字符串的world在最末端的位置
//	l := strings.LastIndex(n, "world")
//	fmt.Println("获取world的最末端的位置：", l)
//
//	//截取m往后的字符串
//	k := n[m:]
//	fmt.Println("截取m往后的字符串：", k)
//
//	//截取m位置往后的3位字符串
//	p := n[m : m+3]
//	fmt.Println("截取m位置往后的3位字符串：", p)
//}

//func main() {
//	n := "hello-world-你好呀"
//	//对字符串惊醒空格分割
//	m := strings.Split(n, "-")
//	fmt.Printf("分割后的数据类型: %T\n", m)
//	//输出分割后的每个字符串
//	for _, i := range m {
//		fmt.Println(i)
//	}
//}

func main() {
	s := "hello world I am Tom"
	//参数n代表替换的次数，从左边开始计算，-1代表替换全部
	m := strings.Replace(s, " ", "-", -1)
	fmt.Println(m)
	//参数n等于1代表只替换一次
	k := strings.Replace(s, " ", "-", 1)
	fmt.Println(k)
}
