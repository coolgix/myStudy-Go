package main

import "fmt"

//func myfun(name string, age int) (string, bool) {
//
//	//参数 name和age
//	//string,bool 返回值的数据类型
//	var n string
//	var b bool
//	if name != "" {
//		//字符串拼接
//		n = name + " is existence, age is  " + strconv.Itoa(age)
//		b = true
//	} else {
//		name = "name is exstence"
//		b = false
//	}
//	//返回值
//	return n, b
//
//}
//
//func main() {
//	//调用函数，并设置返回值
//	s, _ := myfun("tom", 15)
//	fmt.Println(s)
//	// 调用函数，虽然有返回值，但函数外不需要使用
//	myfun("Tom", 15)

//func myfun(numbers ...int) {
//	for _, k := range numbers {
//		fmt.Printf("参数值为：%v\n", k)
//	}
//}
//
//func main() {
//	//调用函数
//	myfun(15, 1231, 412, 41)
//	myfun(15, 11, 41, 4)
//
//}

//func myfun(number ...interface{}) {
//	for _, k := range number {
//		fmt.Printf("参数值为：%v\n", k)
//	}
//}
//
//func main() {
//	var s = []string{"Mary", "Tim"}
//	var m = map[string]interface{}{"name": "Mary", "age": 10}
//	//调用函数
//	myfun(45, "Tom", s, m)
//}

//func myfun() {
//	//定义函数
//	fmt.Printf("自定义函数")
//}
//
//func main() {
//	//定义函数变量
//	var m func()
//	//将函数作为变量m的值
//	m = myfun
//	//调用函数
//	m()
//}

////匿名函数
//func main() {
//	res := func(n1 int, n2 int) int {
//		return n1 + n2
//	}(10, 30)
//	fmt.Printf("函数执行结果为：%v\n", res)
//}

//func main() {
//	// 将匿名函数赋给函数变量myfun
//	myfun := func(n1 int, n2 int) int {
//		return n1 - n2
//	}
//	// 变量myfun的数据类型是函数类型，可以由该变量完成函数调用
//	res2 := myfun(10, 30)
//	res3 := myfun(50, 30)
//	fmt.Printf("匿名函数调用第一次：%v\n", res2)
//	fmt.Printf("匿名函数调用第二次：%v\n", res3)
//	fmt.Printf("函数变量myfun的数据类型：%T\n", myfun)
//}

////闭包
//// 闭包 = 函数 + 引用环境
//func adder() func(int) int {
//	// 定义函数adder()，返回值为匿名函数func(int) int
//	var x int = 10
//	//匿名函数作为函数返回值
//	return func(y int) int {
//		x += y
//		return x
//	}
//}
//func main() {
//	// 函数adder()是一个闭包:
//	// 函数adder()内部有变量x（引用环境）和匿名函数
//	// 匿名函数引用了其外部作用域中的变量x
//	// 在函数adder()的生命周期内 变量x一直有效
//	f := adder()
//	fmt.Println(f(10))
//	fmt.Println(f(20))
//	fmt.Println(f(30))
//	f1 := adder()
//	fmt.Println(f1(2000))
//	fmt.Println(f1(5000))
//	fmt.Println(f1(6000))
//}

func fibonacci(n int) int {
	//定义递归函数
	if n < 2 {
		return n
	}
	//调用自身，传入不同的参数值
	return fibonacci(n-2) + fibonacci(n-1)
}
func main() {
	var i int
	//调用函数fibonacci()
	for i = 0; i < 10; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
}
