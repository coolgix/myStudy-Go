package main

import "fmt"

func main() {
	var n int = 10
	//关系运算符
	r := n == 20
	fmt.Printf("10是否等于20：%v\n", r)

	//if 语句
	if n > 0 {
		fmt.Printf("if的判断条件为true\n")

	}
	//for 语句
	for n > 0 {
		fmt.Printf("for的判断条件为true\n")
		break
	}

}