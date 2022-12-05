package main

import "fmt"

/*
定义一个test函数,他会返回两个int类型的值,每次调用会返回100和200两个数值
*/
func test() (int, int) {
	return 100, 200
}
func main() {
	a, _ := test()
	//这里需要获取第一个返回值,所以第二个返回值定义为匿名变量
	_, b := test()
	//这里需要获取第二个返回值,所以第一个返回值定义为匿名变量
	fmt.Println(a, b)
}
