package main

import (
	"fmt"
	"time"
)

func main() {
	//获取当前时间本地时间（CST）
	now := time.Now()
	fmt.Printf("当前本地(CST)时间：%v\n", now)
	//将当前时间转换为UTC时间
	now1 := now.UTC()
	fmt.Printf("当前本地时间(UTC)：%v\n", now1)
	// 获取当前时间
	now2 := now.UTC()
	fmt.Printf("当前本地时间(UTC)：%v\n", now2)
	// 判断本地(CST)时间和UTC时间是否相同
	r := now.Equal(now1)
	fmt.Printf("判断本地(CST)时间和UTC时间是否相同:%v\n", r)
	//判断一个时间是否存在另一个时间之前
	r1 := now.Before(now2)
	fmt.Printf("判断一个时间是否存在另一个时间之前:%v\n", r1)
	//判断一个时间是否存在另一个时间之后
	r2 := now.After(now2)
	fmt.Printf("判断一个时间是否存在另一个时间之后:%v\n", r2)

}
