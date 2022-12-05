package main

import "fmt"

func main() {
	var a int = 3
	var b int = 7
	//声明局部变量a,b并赋值
	c := a + b
	//声明局部变量c并计算a+b的值
	fmt.Printf("a=%d,b=%d,c=%d", a, b, c)
}
