package main

import (
	"fmt"
)

func main() {
	// var 变量名 数据类型
	// bool默认值是flase
	var isFlag bool = true
	fmt.Println(isFlag)

	fmt.Printf("%T,%t \n", isFlag, isFlag)
}
