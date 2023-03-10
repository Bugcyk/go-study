package main

import (
	"fmt"
)

func main() {
	// 60   0011 1100
	// 13   0000 1101
	// --------------
	// &    0000 1100
	// |    0011 1101
	// ^    0011 0001
	// &^   0011 0000
	// <<2  1111 0000
	// >>2  0000 1111
	var a uint = 60
	var b uint = 13
	var c uint
	fmt.Printf("%d,二进制:%b\n", a, a)
	fmt.Printf("%d,二进制:%b\n", b, b)
	c = a & b
	fmt.Printf("%d,二进制:%b\n", c, c)
	c = a | b
	fmt.Printf("%d,二进制:%b\n", c, c)
	c = a ^ b
	fmt.Printf("%d,二进制:%b\n", c, c)
	c = a &^ b
	fmt.Printf("%d,二进制:%b\n", c, c)
	c = a << 2
	fmt.Printf("%d,二进制:%b\n", c, c)
	c = a >> 2
	fmt.Printf("%d,二进制:%b\n", c, c)
}
