package main

import "fmt"

func main() {
	/*
		var name string = "go study"
		var age int = 18
		fmt.Println(name)
	*/
	var (
		name string // string类型不予赋值情况下,会有一个默认值,null
		age  int    // 默认值为0
		addr string
	)
	//var a, b, c int // 可以将多个变量都声明为一种类型
	name = "go"
	age = 10
	addr = "zhongguo"
	fmt.Println(name, age, addr)
}
