package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(b / a)
	fmt.Println(b % a)
	a++
	fmt.Println(a)
	b--
	fmt.Println(b)
}
