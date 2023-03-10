package main

import "fmt"

func main() {
	var str string = "你好"
	fmt.Printf("%T,%s\n", str, str)

	// 单引号 字符 整型-ASCII字符码
	s1 := 'A'
	// 编码表 ASCII
	// 拓展:
	// 所有中国字的编码表: GBK
	// 全世界的编码表: Unicode编码表
	s2 := "A"
	s3 := '蔡'
	fmt.Printf("%T,%d\n", s1, s1) // int32,65
	fmt.Printf("%T,%s\n", s2, s2)
	fmt.Printf("%T,%d\n", s3, s3)

	// 字符串连接
	fmt.Println("你好" + ",go")

	// 转义字符
	fmt.Println("你好\"go")
	fmt.Println("你好\ngo") // \n 换行
	fmt.Println("你好\tgo") // \t 制表符
}
