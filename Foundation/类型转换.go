package main

import "fmt"

func main() {
	// 数字类型不能转换bool类型
	a := 1
	b := float64(a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T", b)
}
