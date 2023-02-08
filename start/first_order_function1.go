package main

import (
	"fmt"
	"math/rand"
)

/**
在Go里，函数是头等函数，他可以用在整数，字符串或其他类型能用的地方
- 将函数赋给变量
- 将函数作为参数传递
- 将函数作为函数的返回类型
*/

type kelvin2 float64

func fakeSensor() kelvin2 {
	return kelvin2(rand.Intn(151) + 150)
}

func realSensor() kelvin2 {
	return 0
}

func main() {
	// 将函数赋给变量，如果函数没有加()，则不执行，加上()才会执行
	// 变量sensor就是一个函数，而不是函数执行的结果
	// 无论sensor的值是fakeSensor还是realSensor，都可以通过sensor()来调用
	// sensor这个变量的类型是函数，该函数没有参数，返回一个kelvin2类型的值。
	// 换一种声明形式的话：var sensor func() kelvin
	sensor := fakeSensor
	fmt.Println(sensor())

	sensor = realSensor
	fmt.Println(sensor())
}
