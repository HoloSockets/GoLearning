package main

import "fmt"

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
var name, age = "HoloSockets", 18

//全局变量
var m = 18

func main() {

	//局部变量
	n := 20
	m := 19
	//fmt.Println(name)
	//fmt.Println(age)
	fmt.Println(m, n)
}
