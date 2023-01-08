package main

import (
	"fmt"
	"math"
)

func float1() {
	// 这三者是一样的，只要数字含有小数部分，那么它的类型就是float64
	days1 := 365.2425
	var days2 = 365.2425
	var days3 float64 = 365.2425

	fmt.Println(days1, days2, days3)

	// 如果你使用一个整数来初始化某个变量，那么你必须指定它的类型为float64，否则他就是一个整数类型：
	// float64是双精度，float32是单精度
	// Go语言中默认就是64位浮点类型，占用8字节内存；32位浮点类型占用4字节内存
	var answer float64 = 42
	fmt.Println(answer)

	var pi64 = math.Pi
	var pi32 float32 = math.Pi
	fmt.Println(pi64)
	fmt.Println(pi32)
}

// 对于精度要求比较高的情况下，不建议使用浮点数来计算
func float2() {
	third := 1.0 / 3

	fmt.Println(third)
	fmt.Printf("%v\n", third)
	fmt.Printf("%f\n", third)
	fmt.Printf("%.3f\n", third)
	fmt.Printf("%4.2f\n", third)

	third2 := 1.0 / 3.0
	fmt.Println(third2 + third2 + third2)

	piggyBank := 0.1
	piggyBank += 0.2
	fmt.Println(piggyBank)
}

func main() {
	float1()
	fmt.Println("----------")
	float2()
}
