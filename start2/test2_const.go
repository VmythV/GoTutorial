package main

import "fmt"

//const 来定义枚举类型
const (
	// 可以在const() 添加关键字 iota，每行的iota都会累加1，第一行的iota的默认值是0
	BEIJING = iota
	SHANGHAI
	GUANGZHOU
	SHENZHEN
)

const (
	a, b = iota + 1, iota + 2 // iota = 0, a = iota + 1, b = iota + 2, a = 1, b = 2
	c, d                      // iota = 1, c = iota + 1, d = iota + 2, c = 2, d = 3
	e, f                      // iota = 2, e = iota + 1, f = iota + 2, e = 3, f = 4

	g, h = iota * 2, iota * 3 // iota = 3, g = iota * 2, h = iota * 3, g = 6, h = 9
	i, k                      // iota = 4, i = iota * 2, k = iota * 3, i = 8, k = 12
)

func main() {

	const length int = 10

	fmt.Println("length = ", length)

	// length = 100 // 常量是不允许修改的。

	// iota 只能够配合const() 一起使用，iota只有在const进行累加效果
	// var a int = iota

}
