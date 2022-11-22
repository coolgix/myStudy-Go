package main

import "fmt"

func main() {
	var f32 float32 = 1.1
	var f64 float64 = 2.2

	//将float32转换为64
	r := float64(f32) + f64
	fmt.Printf("运行结果为：%v\n", r)

	//将整形转换为浮点数进行运算
	var i int = 10
	rd := float32(i) + f32
	fmt.Printf("运行结果为：%v\n", rd)

}
