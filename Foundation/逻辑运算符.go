package main

import "fmt"

func main() {
	var a bool = true
	var b bool = false

	// && 与;都要满足
	if a && b {
		fmt.Println("true1")
	} else {
		fmt.Println(a && b)
	}

	// || 或;一个满足
	if a || b {
		fmt.Println("true1")
	} else {
		fmt.Println(a || b)
	}

	// ! 非;取反
	if !b {
		fmt.Println("false1")
	} else {
		fmt.Println(!b)
	}
}
