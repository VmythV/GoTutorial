package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
在Go里，函数是头等函数，他可以用在整数，字符串或其他类型能用的地方
- 将函数赋给变量
- 将函数作为参数传递
- 将函数作为函数的返回类型
*/

type kelvin3 float64

// 因为可以声明函数类型，所以可以用以下的方式精简和明确调用者的代码
// type sensor func() kelvin
// func measureTemperature(samples int, s sensor)
func measureTemperature(samples int, sensor func() kelvin3) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%v K\n", k)
		time.Sleep(time.Second)
	}
}

func fakeSensor3() kelvin3 {
	return kelvin3(rand.Intn(151) + 150)
}

func main() {
	measureTemperature(3, fakeSensor3)
}
