package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c
		d = "go"
		e
		f = 100
		g
		h = iota // 行数计数器此时为7
		i
	)
	const (
		j = iota //一组新的变量,计数重新开始
		k
	)
	fmt.Println(a, b, c, d, e, f, g, h, i, j, k)
}
