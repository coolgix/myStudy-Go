package main

import "fmt"

func main() {
	//定义切片的三种方式
	//var s []int
	//var ss = []int{1, 2}
	//var sss []int = make([]int, 3)
	//fmt.Printf("只定义：%v，内存地址为：%v\n", s, &s)
	//fmt.Printf("定义并赋值：%v，内存地址为：%v\n", ss, &ss)
	//fmt.Printf("使用make()函数定义：%v，内存地址为：%v\n", sss, &sss)

	////修改切片的元素
	//var ss = []int{1, 2}
	//fmt.Printf("切片变量ss的元素值为：%v\n", ss)
	//
	////修改第一个元素.修改第一个元素的索引
	//ss[0] = 100
	//fmt.Printf("切片变量ss的元素值为：%v\n", ss)
	//
	////修改元素的索引值大于切片的长度，则报错
	//ss[10] = 100
	//fmt.Printf("切片变量ss的元素值为：%v\n", ss)

	////新增切片元素
	//var ss = []int{1, 2}
	//fmt.Printf("新增元素前的切片ss：%v\n", ss)
	//
	////新增元素不覆盖原有的切片
	//sss := append(ss, 3)
	//fmt.Printf("新增元素后的切片ss：%v\n", ss)
	//fmt.Printf("新切片sss：%v\n", sss)
	//
	////新增元素并覆盖原有切片
	//ss = append(ss, 3)
	//fmt.Printf("新增元素后的切片ss：%v\n", ss)
	//
	////添加多个元素
	//ss = append(ss, 5, 6, 7, 8, 9, 0)
	//fmt.Printf("新增元素后的切片ss：%v\n", ss)

	////内置函数append 实现两个切片的拼接
	//var s1 = []int{1, 2, 3}
	//var s2 = []int{4, 5, 6, 7}
	//ss := append(s1, s2...)
	//fmt.Printf("切片变量s1:%v\n", s1)
	//fmt.Printf("切片变量s2:%v\n", s2)
	//fmt.Printf("切片变量ss:%v\n", ss

	////截取切片元素
	//var ss = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//
	////截取第二个到第五个元素
	//s1 := ss[1:4]
	//fmt.Printf("截取第二个到第五个元素：%v\n", s1)
	//
	////截取第三个元素之后的所有元素
	//s2 := ss[2:]
	//fmt.Printf("截取第三个元素之后的所有元素：%v\n", s2)
	//
	////截取第三个元素之前的所有元素
	//s3 := ss[:2]
	//fmt.Printf("截取第三个元素之前的所有元素：%v\n", s3)
	//
	//// 如果切片ss没被覆盖，经过截取后不改变原有的切片数据
	//fmt.Printf("切片变量ss的值：%v\n", ss)

	////删除切片的部分元素
	//var ss = []int{1, 2, 3, 4, 5, 6, 7}
	//fmt.Printf("切片ss的元素：%v\n", ss)
	////删除元素456
	////先截取，后用append合并
	//ss = append(ss[:2], ss[6:]...)
	//fmt.Printf("切片ss的元素：%v\n", ss)

	////复制切片
	//slice1 := []int{1, 2, 3}
	//slice2 := []int{4, 5, 6}
	//// 将slice1的元素复制到slice2
	////copy(slice2, slice1)
	////fmt.Printf("将slice2的元素复制到slice1：%v\n", slice1)
	////fmt.Printf("将slice2的元素复制到slice1：%v\n", slice2)
	//
	//// 将slice2的元素复制到slice1
	//copy(slice1, slice2)
	//fmt.Printf("将slice2的元素复制到slice1：%v\n", slice1)
	//fmt.Printf("将slice2的元素复制到slice1：%v\n", slice2)
	//slice3 := []int{7, 8, 9, 10}
	//
	////将slice3的元素复制到slice1
	//copy(slice1, slice3)
	//fmt.Printf("将slice3的元素复制到slice1：%v\n", slice1)
	//fmt.Printf("将slice3的元素复制到slice1：%v\n", slice3)
	//
	//// 将slice1的元素复制到slice3
	//copy(slice3, slice1)
	//fmt.Printf("将slice1的元素复制到slice3：%v\n", slice1)
	//fmt.Printf("将slice1的元素复制到slice3：%v\n", slice3)

	//切片的长度与容量
	// 内置函数cap()获取切片容量
	// 内置函数len()获取切片长度
	s1 := make([]int, 3, 4)
	fmt.Printf("切片变量s1的值：%v\n", s1)
	fmt.Printf("切片变量s1的长度：%v\n", len(s1))
	fmt.Printf("切片变量s1的容量：%v\n", cap(s1))

	// 第一次添加元素
	s1 = append(s1, 10)
	fmt.Printf("切片变量s1的值：%v\n", s1)
	fmt.Printf("切片变量s1的长度：%v\n", len(s1))
	fmt.Printf("切片变量s1的容量：%v\n", cap(s1))

	//第二次添加元素
	s1 = append(s1, 10)
	fmt.Printf("切片变量s1的值：%v\n", s1)
	fmt.Printf("切片变量s1的长度：%v\n", len(s1))
	fmt.Printf("切片变量s1的容量：%v\n", cap(s1))
}
