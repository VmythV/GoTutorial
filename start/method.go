package main

import "fmt"

func main() {
	func1()
	func2()
}

func func1() {
	// 关键字type可以用来声明新类型
	// 为什么要声明新类型：极大的提高代码的可读性和可靠性
	type celsius float64

	const degrees = 20
	var temperature celsius = degrees

	temperature += 10

	fmt.Println(temperature)
}

func func2() {
	var k kelvin = 294.0
	var c celsius

	c = kelvinToCelsius(k)
	c = k.celsius()

	fmt.Println(c)
}

type kelvin float64
type celsius float64

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}
