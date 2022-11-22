package main

import (
	"fmt"
	"strconv"
)

func main() {

	//定义一个空的字符串类型的切片指针
	var pslice []*string

	//定义字符串类型变量a
	//var a string

	//循环5次，当前循环次数赋值给变量a,再次写入切片指针
	for i := 0; i < 5; i++ {
		//定义字符串类型变量a
		//每次循环都是重新定义变量a的值，重新赋予变量a的内存地址
		var a string
		a = strconv.Itoa(i)
		pslice = append(pslice, &a)
	}

	//输出切片指针的元素的数值
	for _, k := range pslice {
		fmt.Printf("切片指针的元素：%v，元素的值：%v\n", k, *k)
	}
}
