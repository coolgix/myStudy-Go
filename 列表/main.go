package main

import (
	"container/list"
	"fmt"
)

func main() {

	//// 定义列表变量
	//var l2 list.List
	//fmt.Printf("列表l2：%v\n", l2)
	//
	//// 在列表最末位位置新增元素，返回当前元素信息
	//l2.PushBack("a")
	//// 在列表最首位位置新增元素，返回当前元素信息
	//l2.PushFront(67)
	//fmt.Printf("列表l2：%v\n", l2)
	//
	//// 定义列表变量
	//l1 := list.New()
	//fmt.Printf("列表l1：%v\n", l1)
	//
	//// 在列表最末位位置新增元素，返回当前元素信息
	//element := l1.PushBack("abc")
	//fmt.Printf("元素element：%v\n", element)
	//
	//// 在元素element（即abc）后面添加元素
	//l1.InsertAfter("edf", element)
	//fmt.Printf("在元素element后面添加元素：%v\n", l1)
	//
	//// 在元素element（即abc）前面添加元素
	//l1.InsertBefore("ghi", element)
	//fmt.Printf("在元素element前面添加元素：%v\n", l1)
	//
	//// 列表l2的元素添加列表l1的元素后面
	//l1.PushBackList(&l2)
	//fmt.Printf("列表l2添加列表l1后面：%v\n", l1)
	//
	//// 将列表l2的元素添加到列表l1的元素前面
	//l1.PushFrontList(&l2)
	//fmt.Printf("列表l2添加列表l1前面：%v\n", l1)

	////定义列表变量
	//var l2 list.List
	//fmt.Printf("列表l2：%v\n", l2)
	//
	////添加元素
	//element := l2.PushBack("abc")
	//fmt.Printf("列表l2：%v\n", l2)
	//
	////删除元素
	//l2.Remove(element)
	//fmt.Printf("列表l2：%v\n", l2)

	////定义列表变量
	//var l2 list.List
	////添加元素
	//l2.PushBack("Tom")
	//l2.PushBack("Tim")
	//l2.PushBack("Mary")
	////遍历输出元素
	//for i := l2.Front(); i != nil; i = i.Next() {
	//	fmt.Printf("列表l2的元素是：%v\n", i.Value)
	//}

	//定义列表变量
	var l2 list.List
	//添加元素
	l2.PushBack("Tom")
	l2.PushBack("Tim")
	l2.PushBack("Lily")
	l2.PushBack("Mary")
	fmt.Printf("列表l2的元素是：%v\n", l2)
	//定义变量next
	var next *list.Element
	//遍历输出元素
	for i := l2.Front(); i != nil; i = next {
		fmt.Printf("列表l2的元素是：%v\n", i.Value)
		//设置变量next的值，用于执行下一次循环
		next = i.Next()
		//删除元素
		l2.Remove(i)

	}
	fmt.Printf("列表l2的元素是：%v\n", l2)
}
