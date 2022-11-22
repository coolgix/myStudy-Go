package main

import "fmt"

//func main() {
//	var x, y = 8, 5
//	//算术运算
//	fmt.Printf("加法运算符：%d\n", x+y)
//	fmt.Printf("减法运算符： %d\n", x-y)
//	fmt.Printf("乘法运算符： %d\n", x*y)
//	fmt.Printf("除法运算符： %d\n", x/y)
//	fmt.Printf("求余运算符： %d\n", x%y)
//	x++
//	fmt.Printf("幂运算符： %d\n", x)
//	y--
//	fmt.Printf("取整运算符：%d\n", y)
//
//	//关系运算
//	fmt.Printf("x大于y：%v\n", x > y)
//	fmt.Printf("x大于或等于y：%v\n", x >= y)
//	fmt.Printf("x小于y：%v\n", x < y)
//	fmt.Printf("x小于或等于y：%v\n", x <= y)
//	fmt.Printf("x等于y：%v\n", x == y)
//	fmt.Printf("x不等于y：%v\n", x != y)
//
//	//赋值运算
//	var a, c int = 21, 0
//	c = a
//	fmt.Printf("=运算符实例，c值为=%d\n", c)
//	c, a = 1, 20
//	c += a
//	fmt.Printf("+=运算符实例，c值为=%d\n", c)
//	c, a = 1, 20
//	c -= a
//	fmt.Printf("-=运算符实例，c值为=%d\n", c)
//	c, a = 1, 20
//	c *= a
//	fmt.Printf("*=运算符实例，c值为=%d\n", c)
//	c, a = 1, 20
//	c /= a
//	fmt.Printf("/=运算符实例，c值为=%d\n", c)
//	c, a = 1, 20
//	c %= a
//	fmt.Printf("求余运算符实例，c值为=%d\n", c)
//	c = 200
//	c <<= 2
//	fmt.Printf("<=运算符实例，c值为=%d\n", c)
//	c = 200
//	c >>= 2
//	fmt.Printf(">=运算符实例，c值为=%d\n", c)
//	c = 200
//	c &= 2
//	fmt.Printf("&=运算符实例，c值为=%d\n", c)
//	c = 200
//	c ^= 2
//	fmt.Printf("^=运算符实例，c值为=%d\n", c)
//	c = 200
//	c |= 2
//	fmt.Printf("|=运算符实例，c值为=%d\n", c)
//
//}

//func main() {
//	//逻辑运算符
//	var a, b, c, d = true, true, false, false
//
//	fmt.Printf("a&&b的值为：%v\n", a && b)
//	fmt.Printf("a&&c的值为：%v\n", a && c)
//	fmt.Printf("a&&d的值为：%v\n", a && d)
//
//	fmt.Printf("a||b的值为：%v\n", a || b)
//	fmt.Printf("a||c的值为：%v\n", a || c)
//	fmt.Printf("a||d的值为：%v\n", a || d)
//
//	fmt.Printf("!a的值为：%v\n", !a)
//	fmt.Printf("!c的值为：%v\n", !c)
//
//}

//func main() {
//	//位运算
//	var a = 60 // 60=0011 1100
//	var b = 13 // 13=0000 1101
//	var c = 0
//
//	c = a & b //12=0000 1100
//	fmt.Printf("c的十进制为：%d\n", c)
//	fmt.Printf("c的二进制为：%b\n", c)
//
//	c = a | b //61=0011 1101
//	fmt.Printf("c的十进制为：%d\n", c)
//	fmt.Printf("c的二进制为：%b\n", c)
//
//	c = a ^ b //49=0011 0001
//	fmt.Printf("c的十进制为：%d\n", c)
//	fmt.Printf("c的二进制为：%b\n", c)
//
//	c = a << 2 //240=1111 0000
//	fmt.Printf("c的十进制为：%d\n", c)
//	fmt.Printf("c的二进制为：%b\n", c)
//
//	c = a >> 2 //15=0000 1111
//	fmt.Printf("c的十进制为：%d\n", c)
//	fmt.Printf("c的二进制为：%b\n", c)
//
//}

func main() {
	var a = 4
	var ptr *int
	fmt.Printf("a变量的内存地址：%v\n", &a)
	//将变量a的内存地址指向给ptr
	ptr = &a
	fmt.Printf("ptr的内存地址：%v\n", ptr)
	fmt.Printf("指针ptr的值为：%v\n", *ptr)

}