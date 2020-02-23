package main

import "fmt"

func f1(ff func()) {
	fmt.Println("this is f1")
	ff()
}

func f2(x, y int) int {
	fmt.Println("this is f2")
	return x + y
}

func f3(tempF2 func(int, int) int, x, y int) func() {
	return func() {
		z := tempF2(x, y)
		fmt.Println(z)
	}
}

func main() {
	f1(f3(f2, 1, 2))
}
