package main

import "fmt"

func main() {
	//var a int = 200
	//fmt.Printf("变量a的空间地址：%v\n", &a)
	//
	////定义int类型的指针变量
	//var pint *int //空指针
	//fmt.Printf("指针值为：%v，空间地址：%v\n", pint, &pint)
	//
	////给指针赋值
	//pint = &a
	//fmt.Printf("指针值为：%v，空间地址：%v\n", pint, &pint)
	////*pint ，取出变量a的值，内存地址不变
	//fmt.Printf("指针值的值为：%v，空间地址：%v\n", *pint, &pint)

	var b int = 200
	var pint *int //空指针 未赋值
	fmt.Printf("指针值的值为：%v，空间地址：%v\n", pint, &pint)

	pint = &b
	fmt.Printf("指针值的值为：%v，空间地址：%v\n", *pint, &pint)

	// 通过取值操作符“*”修改变量b的值,内存地址不变，取值变了
	*pint = 666
	fmt.Printf("指针值的值为：%v，空间地址：%v\n", *pint, &pint)
}
