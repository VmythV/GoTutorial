package main

import "fmt"

// 闭包和匿名函数
// 匿名函数就是没有名称的函数，在Go里也称函数字面值
// 因为函数字面值需要保留外部作用域的变量引用，所以函数字面值都是闭包的

var f = func() {
	fmt.Println("Dress up for the masquerade")
}

func main() {
	// 第一种形式，声明匿名函数
	f()

	// 第二种形式，在main包下声明
	f2 := func(message string) {
		fmt.Println(message)
	}
	f2("xxx")

	// 第三种形式，直接声明完调用
	func() {
		fmt.Println("YYY")
	}()
}
