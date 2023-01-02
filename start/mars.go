package main

import "fmt"

/*
	const 用来声明常量（常量的值不可以改变）
	var   用来声明变量（想要使用变量首先需要进行声明）

	可以同时声明多个变量
*/

func main() {
	fmt.Print("My weight on surface of Mars is ")
	fmt.Print(150 * 0.3783)
	fmt.Print(" lbs, and I would be ")
	fmt.Print(41 * 365 / 687)
	fmt.Print(" year old.\n")

	fmt.Println("My weight on surface of Mars is", 150*0.3783, "lbs, and I would be", 41*365/687, "year old.")

	// 常量
	const lightSpeed = 299792 // km/s
	// 变量
	var distance = 56000000 // km

	fmt.Println(distance/lightSpeed, "seconds")

	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")

	var speed = 100800

	var (
		distance2 = 56000000
		speed2    = 100800
	)

	var distance3, speed3 = 56000000, 100800

	const hoursPerDay, minutesPerHour = 24, 60

	fmt.Println("speed:", speed)
	fmt.Println("distance2:", distance2, "speed2:", speed2, "distance3:", distance3, "speed3:", speed3)

	var weight = 150.0
	weight = weight * 0.3783
	weight *= 0.3783
	fmt.Println("weight:", weight)

	var age = 25
	age = age + 1
	age += 1
	age++
	// 没有 ++age
	fmt.Println("age", age)
}
