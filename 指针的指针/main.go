package main

import "fmt"

func main() {
	var str string = "hello"

	var ptr *string

	var pptr **string

	ptr = &str
	pptr = &ptr

	fmt.Printf("字符串str为：%v，空间地址为：%v\n", str, &str) //hello xxxx01

	fmt.Printf("指针变量ptr为：%v，空间地址为：%v\n", ptr, &ptr) //xx01 xxxx02

	fmt.Printf("指针的指针pptr为：%v，空间地址为：%v\n", pptr, &pptr) //xxx02 xxxx03

	//从指针的指针取某个变量值
	fmt.Printf("指针的指针pptr取变量str的值：%v\n", **pptr) //hello
}
