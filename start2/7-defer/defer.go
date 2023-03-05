package main

import "fmt"

func main() {
	// 写入defer关键字，defer执行顺序是按照压栈的顺序执行的，也就是说越靠前的defer越后执行
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")
}
