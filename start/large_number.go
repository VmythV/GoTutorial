package main

import (
	"fmt"
	"math/big"
)

// Go中提供大数的包big，对于比较大的整数（超过10^18）:big.Int，对于任意精度的浮点数类型：big.Float，对于分数：big.Rat
func main() {
	const lightSpeed = 299792
	const secondsPerDay = 86400

	var distance int64 = 41.3e12
	fmt.Println("Alpha Centauri is", distance, "km away.")

	days := distance / lightSpeed / secondsPerDay
	fmt.Println("That is", days, "days of travel at light speed.")

	fmt.Println("---------------------------------")

	lightSpeed2 := big.NewInt(299792)
	secondsPerDay2 := big.NewInt(86400)
	fmt.Println(lightSpeed2, secondsPerDay2)

	fmt.Println("---------------------------------")
	// new 一个实例，实际上是一个指针
	distance3 := new(big.Int)
	// 设置值，前面是字符串的数字，后面是进制
	// 实例有多个方法，distance3.Div()：除法
	distance3.SetString("2400000000000000000000000000000", 10)
	fmt.Println(distance3)

	fmt.Println("---------------------------------")
	// 在Go里面，可以为常量指明类型，也可以不指定类型，对于变量，Go会使用类型推断
	// 在Go中，常量是可以无类型的(untyped)
	// 针对字面值和常量的计算是在编译阶段完成的

}
