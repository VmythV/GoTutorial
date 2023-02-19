package main

import (
	"fmt"
)

// 在面向对象的世界中，对象是由更小的对象组合而成
// Go中通过结构体实现组合（composition）
// Go提供了"嵌入"（embedding）特性，它可以实现方法的转发（forwarding）
// 组合是一种更简单，灵活的方式

type report struct {
	sol         int
	temperature temperature
	location    location2
}

type temperature struct {
	high, low celsius2
}

type location2 struct {
	lat, long float64
}

type celsius2 float64

func (t temperature) average() celsius2 {
	return (t.high + t.low) / 2
}

func (r report) average() celsius2 {
	return r.temperature.average()
}

func main() {
	bradbury := location2{-4.5895, 137.4417}
	t := temperature{high: -1.0, low: -78.0}

	report := report{
		sol:         15,
		temperature: t,
		location:    bradbury,
	}
	fmt.Println(report.temperature.average())
	fmt.Println(report.average())

	fmt.Printf("%+v\n", report)
	fmt.Printf("a balmy %v C\n", report.temperature.high)
}
