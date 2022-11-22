package main

import "fmt"

//定义结构体person

//type person struct {
//	name string
//	age  int
//}

//4种实例化方式
//func main() {
//	//实例化方法1
//	//实例化结构体person，生成实例化对象p
//	p := person{
//		name: "Tom",
//		age:  18,
//	}
//	//由实例化对象p访问成员
//	fmt.Printf("结构体成员name的值：%v\n", p.name)
//	fmt.Printf("结构体成员age的值：%v\n", p.age)
//
//	//实例化方法2
//	//实例化结构体
//	var p1 person
//	//对结构体成员进行赋值操作
//	p1.name = "Tim"
//	p1.age = 22
//	// 由实例化对象p1访问成员
//	fmt.Printf("结构体成员name的值：%v\n", p1.name)
//	fmt.Printf("结构体成员age的值：%v\n", p1.age)
//
//	// 实例化方法3
//	// 使用new()实例化结构体
//	p3 := new(person)
//	// 对结构体成员进行赋值操作
//	p3.name = "LiLy"
//	p3.age = 28
//	// 由实例化对象p3访问成员
//	fmt.Printf("结构体成员name的值：%v\n", p3.name)
//	fmt.Printf("结构体成员age的值：%v\n", p3.age)
//
//	// 实例化方法4
//	// 取结构体的地址实例化
//	p4 := &person{}
//	// 对结构体成员进行赋值操作
//	p4.name = "Mary"
//	p4.age = 16
//	// 由实例化对象p4访问成员
//	fmt.Printf("结构体成员name的值：%v\n", p4.name)
//	fmt.Printf("结构体成员age的值：%v\n", p4.age)
//}

////指针实例化方式
//func main() {
//	// 实例化方法3
//	// 使用new()实例化结构体
//	var p3 *person = new(person)
//	//对结构体成员进行赋值操作
//	(*p3).name = "lily"
//	(*p3).age = 28
//	// 由实例化对象p3访问成员
//	fmt.Printf("结构体成员name的值：%v\n", p3.name)
//	fmt.Printf("结构体成员age的值：%v\n", p3.age)
//
//	// 实例化方法4
//	// 取结构体的地址实例化
//	var p4 *person = &person{}
//	//对结构体成员进行赋值操作
//	(*p4).name = "Mary"
//	(*p4).age = 16
//	// 由实例化对象p4访问成员
//	fmt.Printf("结构体成员name的值：%v\n", p4.name)
//	fmt.Printf("结构体成员age的值：%v\n", p4.age)
//}

////结构体标签
//type Student struct {
//	name  string
//	age   int
//	score int
//}
////小写的首字母不是json包里认为的导出标识符，json包无法成功获取到结构体里面每个成员的数据
//func main() {
//	var stu Student = Student{
//		name:  "张三",
//		age:   22,
//		score: 88,
//	}
//	data, _ := json.Marshal(stu)
//	fmt.Printf("结构体转换JSON：%v\n", string(data))
//}

//type Student struct {
//	Name  string
//	Age   int
//	Score int
//}
////应为json里面的数据的key一般是小写
//func main() {
//	var stu Student = Student{
//		Name:  "张三",
//		Age:   22,
//		Score: 88,
//	}
//
//	data, _ := json.Marshal(stu)
//	fmt.Printf("结构体转换JSON：%v\n", string(data))
//}

////应为json里面的数据的key一般是小写
////通过tag实现数据类型转换
//type Student struct {
//	// `json:"name"`表示json序列化时，结构体成员展示形式为name
//	Name  string `json:"name"`
//	Age   int    `json:"age"`
//	Score int    `json:"score"`
//}
//
//func main() {
//	var stu Student = Student{
//		Name:  "张三",
//		Age:   22,
//		Score: 88,
//	}
//
//	data, _ := json.Marshal(stu)
//	fmt.Println(string(data))
//}

//匿名结构体与匿名成员
//func main() {
//	var p struct {
//		name string
//		age  int
//	}
//	//使用匿名结构体并为成员赋值
//	p.name = "Tom"
//	p.age = 10
//	fmt.Printf("匿名结构体的成员name：%v\n", p.name)
//	fmt.Printf("匿名结构体的成员age：%v\n", p.age)
//
//	// 定义匿名结构体并赋值
//	p1 := struct {
//		name string
//		age  int
//	}{
//		name: "Tim",
//		age:  20,
//	}
//	fmt.Printf("匿名结构体的成员name：%v\n", p1.name)
//	fmt.Printf("匿名结构体的成员age：%v\n", p1.age)
//}

////匿名成员
//type person struct {
//	//定义结构体
//	string
//	int
//	float64
//	bool
//}
//
//func main() {
//	//实例化结构体
//	p := person{"Tim", 10, 171.1, true}
//	// 访问匿名成员并输出
//	fmt.Printf("结构体的匿名成员string的值：%v\n", p.string)
//	fmt.Printf("结构体的匿名成员int的值：%v\n", p.int)
//	fmt.Printf("结构体的匿名成员float64的值：%v\n", p.float64)
//	fmt.Printf("结构体的匿名成员bool的值：%v\n", p.bool)
//}

//结构体嵌套

//func main() {
//	type cars struct {
//		name  string
//		price int
//	}
//
//	type person struct {
//		name   string
//		age    int
//		cars   cars
//		hourse struct { //不符合代码复用的思想所以一般不这杨写
//			name  string
//			price int
//		}
//	}
//
//	c := cars{
//		name:  "BMW",
//		price: 5000000,
//	}
//	p := person{
//		name: "Tim",
//		age:  30,
//		cars: c,
//	}
//	fmt.Printf("个人名称：%v\n", p.name)
//	fmt.Printf("个人年龄：%v\n", p.age)
//	fmt.Printf("个人拥有车辆：%v\n", p.cars.name)
//	fmt.Printf("车辆价钱为：%v\n", p.cars.price)
//}

//自定义构造函数
//type cat struct {
//	//定义结构体
//	name   string
//	weight int
//	//结构体成员为匿名结构体
//	habit struct {
//		ambient string
//		style   string
//	}
//}
//
//func get_cat() *cat {
//	//定义构造函数
//	//设置随机数的种子
//	rand.Seed(time.Now().UnixNano())
//	n := rand.Intn(10)
//	fmt.Printf("随机数：%v\n", n)
//	//定义变量，用于设置结构体的成员值
//	var name, ambient, style string
//	var weight int
//	//根据随机数设置变量值
//	if n <= 5 {
//		name = "tiger"
//		weight = 500
//		ambient = "山林"
//		style = "独居"
//	} else {
//		name = "lion"
//		weight = 300
//		ambient = "草原"
//		style = "群居"
//	}
//	//实例化结构体
//	c := cat{
//		name:   name,
//		weight: weight,
//		//匿名结构体实例化
//		habit: struct {
//			ambient string
//			style   string
//		}{ambient: ambient, style: style},
//	}
//	return &c
//}
//
//func main() {
//	// 调用构造函数，获取结构体实例化对象
//	c := get_cat()
//	fmt.Printf("猫科动物为：%v\n", c.name)
//	fmt.Printf("体重为：%v\n", c.weight)
//	fmt.Printf("居住环境：%v\n", c.habit.ambient)
//	fmt.Printf("生活方式：%v\n", c.habit.style)
//}

//结构体方法：指针与值的接收者

type person struct {
	//定义结构体
	name   string
	weight int
}

func (p person) get_name(name string) string {
	//定义结构体的方法
	return "My name is " + name
}

func (p *person) init(name string, weight int) {
	//定义结构体的方法，用于初始化结构体成员
	p.name = name
	p.weight = weight
}

func main() {
	//实例化结构体
	p := person{}
	//调用结构体方法，初始化成员值
	p.init("Tom", 99)
	fmt.Printf("结构体的成员name的值：%v\n", p.name)
	fmt.Printf("结构体的成员weight的值：%v\n", p.weight)
	name := p.get_name(p.name)
	fmt.Printf("结构体方法的返回值：%v\n", name)
}
