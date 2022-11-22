package main

import (
	"fmt"
	"strconv"
)

func main() {
	//整形转换字符串
	s := strconv.Itoa(100)
	fmt.Printf("整形转为字符串：%T：%v\n", s, s)

	//字符串转换为整形
	i, _ := strconv.Atoi("110")
	fmt.Printf("字符串转为整形：%T：%v", i, i)

	//字符串转为整形int8
	k, _ := strconv.ParseInt("120", 10, 8)
	fmt.Printf("字符串转为整形int8：%T：%v\n", k, k)

	//字符串转为整形int16
	l, _ := strconv.ParseInt("120", 10, 16)
	fmt.Printf("字符串转为整形int16：%T:%v\n", l, l)

	//字符串转为整形int32
	m, _ := strconv.ParseInt("120", 10, 32)
	fmt.Printf("字符串转为整形int32：%T：%v\n", m, m)

	//字符串转为整形int64
	n, _ := strconv.ParseInt("120", 10, 64)
	fmt.Printf("字符串转换整形int64：%T：%v", n, n)
}
