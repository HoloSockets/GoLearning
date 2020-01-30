package main

import "fmt"

//----------------------变量----------------------
//标准声明
// var name string
// var age int
// var isOk bool

//批量声明
// var (
// 	a string
// 	b int
// 	c bool
// )

//变量的初始化
//var name = "HoloSockets" //var name string = "HoloSockets"(一般不这么写，使用前面的类型推导格式)
//var age = 18             //var age int = 18(一般不这么写，使用前面的类型推导格式)

//一次初始化多个变量
//var name, age = "HoloSockets", 18

//全局变量
//var m = 18

//匿名变量
// func foo() (int, string) {
// 	return 10, "HoloSockets"
// }

//变量注意事项：

// 1.函数外的每个语句都必须以关键字开始（var、const、func等）
// 2.:=不能使用在函数外。
// 3._多用于占位，表示忽略值。

//----------------------常量----------------------
//标准声明
// const pi = 3.1415
// const e = 2.7182

//批量声明
// const (
// 	pi = 3.1415
// 	e  = 2.7182
// )

//同时批量声明多个常量，省略了值则和上一行的值相同
// const (
// 	a = 100
// 	b
// 	c
// )

//iota(go语言的常量计数器)
// const (
// 	n1 = iota //0
// 	n2        //1
// 	n3        //2
// 	n4        //3
// )

//使用_跳过某些值
// const (
// 	n1 = iota //0
// 	n2        //1
// 	_         //2
// 	n4        //3
// )

//iota声明中间插队
const (
	n1 = iota //0
	n2 = 100  //100
	n3 = iota //2    n3  //100
	n4        //3    n4  //100
)

const n5 = iota //0

//定义数量级 （这里的<<表示左移操作，1<<10表示将1的二进制表示向左移10位，也就是由1变成了10000000000，也就是十进制的1024。同理2<<2表示将2的二进制表示向左移2位，也就是由10变成了1000，也就是十进制的8。）
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

//多个iota定义在一行
const (
	a, b = 1 + iota, 2 + iota //1 2
	c, d                      //2,3
	e, f                      //3,4
)

func main() {

	//局部变量
	//n := 20
	//m := 19
	//fmt.Println(name)
	//fmt.Println(age)
	//fmt.Println(m, n)
	// age, _ := foo()
	// _, name := foo()
	// fmt.Println("age:", age)
	// fmt.Println("name:", name)
	// fmt.Println("a:", a)
	// fmt.Println("b:", b)
	// fmt.Println("c:", c)
	// fmt.Println("n1:", n1)
	// fmt.Println("n2:", n2)
	// fmt.Println("n3:", n3)
	// fmt.Println("n4:", n4)
	// fmt.Println("n5:", n5)
	// fmt.Println("KB:", KB)
	// fmt.Println("MB:", MB)
	// fmt.Println("GB:", GB)
	// fmt.Println("TB:", TB)
	// fmt.Println("PB:", PB)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
	fmt.Println("f:", f)
}
