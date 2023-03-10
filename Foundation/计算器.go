package main

import "fmt"

func main() {
	var a, b int
	var c string
	fmt.Println("请输入:")
	fmt.Scanln(&a, &c, &b)
	if c == "+" {
		fmt.Println(a + b)
	} else if c == "-" {
		fmt.Println(a - b)
	} else if c == "*" {
		fmt.Println(a * b)
	} else if c == "/" {
		fmt.Println(a / b)
	} else {
		fmt.Println("输入错误,请重新输入!")
	}
}
