package main

import "fmt"

func main() {
	var a byte
	var b rune
	a = 'u'
	b = 'a'
	c := 'c'

	//格式化%T输出变量的数据类型
	fmt.Printf("变量a的数据类型：%T\n", a)
	fmt.Printf("变量b的数据类型：%T\n", b)
	fmt.Printf("变量c的数据类型：%T\n", c)
}
