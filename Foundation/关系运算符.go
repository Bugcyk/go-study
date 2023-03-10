package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)

	// if语句
	if a < b {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
