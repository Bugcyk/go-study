package main

import "fmt"

func main() {
	const URL string = "www.baidu.com"     //显式类型
	const URL1 = "www.baidu.com"           //隐式类型
	const a, b, c = 3.1415926, true, "yes" //多重赋值
	fmt.Println(URL)
	fmt.Println(URL1)
	fmt.Println(a, b, c)
}
