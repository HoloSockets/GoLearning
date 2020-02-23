package main

import (
	"fmt"
	"math"
)

//go语言中的基本数据类型有：整型、浮点型、布尔型、字符串

//-------------------整型-------------------
//按长度分为：int8、int16、int32、int64
//对应的无符号整型：uint8、uint16、uint32、uint64
//其中，uint8(byte型)，int16(short型)，int64(long型)

//-------------------浮点型-------------------
//两种浮点型数：float32和float64

//-------------------复数-------------------
//complex64和complex128

//-------------------布尔值-------------------
//bool,布尔型数据只有true（真）和false（假）两个值
//注意：
// 1.布尔类型变量的默认值为false。
// 2.Go 语言中不允许将整型强制转换为布尔型.
// 3.布尔型无法参与数值运算，也无法与其他类型进行转换。

//-------------------字符串-------------------
//字符串转义符 \r \n \t \' \" \\
//多行字符串 ``(反引号)
//字符串的常用操作

//-------------------byte和rune类型-------------------
//Go 语言的字符有以下两种：
// 1.uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
// 2.rune类型，代表一个 UTF-8字符。

func main() {
	// 十进制
	var a = 10
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) // 1010  占位符%b表示二进制

	// 八进制  以0开头
	var b = 077
	fmt.Printf("%o \n", b) // 77

	// 十六进制  以0x开头
	var c = 0xff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF

	//浮点数
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

	//复数
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)
}
