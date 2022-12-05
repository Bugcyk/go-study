package main

import "fmt"

var c int

// 定义一个全局变量
func main() {
	var a, b int
	//定义局部变量
	a = 3
	b = 7
	//初始化参数
	c = a + b
	fmt.Printf("a=%d,b=%d,c=%d", a, b, c)
}
