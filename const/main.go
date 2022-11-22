package main

import "fmt"

////定义单个常量
//const name type = value
//
////定义单个常量，可以省略常量类型
//const name = value
//
////使用小括号定义多个常量

func main() {
	const (
		a = iota //iota为0。，常量a的值为iota的值
		b        // iota累加1 常量b的值为iota的值
		c = 10   //iota累加1 常量c的值为10
		d        //iota累加1 常量d的值为10
		e = iota //iota 累加1 ，常量e 的值为iota的值

	)
	fmt.Printf("a的值为：%d\n", a)
	fmt.Printf("b的值为：%d\n", b)
	fmt.Printf("c的值为：%d\n", c)
	fmt.Printf("d的值为：%d\n", d)
	fmt.Printf("e的值为：%d\n", e)

	//实现数据存储转换
	const (
		_          = iota             //忽略iota第一个值
		KB float64 = 1 << (10 * iota) //1<< (10*1)
		MB                            //1<< (10*2)
		GB                            //1<< (10*3)
		TB                            //1<< (10*4)

	)

	fmt.Printf("B转KB的进制为 %.0f\n", KB)
	fmt.Printf("B转MB的进制为 %.0f\n", MB)
	fmt.Printf("B转GB的进制为 %.0f\n", GB)
	fmt.Printf("B转TB的进制为 %.0f\n", TB)

}
