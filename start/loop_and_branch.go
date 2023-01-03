package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// 判断
	fmt.Println("You find yourself in a dimly lit cavern. ")

	var command = "walk outside"
	// strings包的Contains函数可以判断某个字符串是否包含另一个字符串
	var exit = strings.Contains(command, "outside")

	fmt.Println("You leave the cave:", exit)

	fmt.Println("==============")

	// 比较
	// Go中比较两个值，得到的结果也是true或false
	// 比较运算符：==，<=，<，!=，>=，>
	fmt.Println("There is a sign near the entrance that reads 'No minors'.")

	var age = 41
	var minor = age < 18

	fmt.Printf("At age %v, am I a minor? %v\n", age, minor)

	fmt.Println("==============")
	// 分支
	var command2 = "go east"
	if command == "go east" {
		fmt.Println("You head further up the mountain.")
	} else if command2 == "go inside" {
		fmt.Println("You enter thee cafe where you live out the test of your life.")
	} else {
		fmt.Println("Didn't quite get that. ")
	}
	fmt.Println("==============")
	// 与，或，取反
	fmt.Println("The year is 2100, should you leap?")

	var year = 2100
	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)

	if leap {
		fmt.Println("Look before you leap!")
	} else {
		fmt.Println("Keep your feet on the ground.")
	}

	fmt.Println("==============")
	// switch
	// switch语句可以对数字进行匹配
	// Go的switch还有一个fallthrough关键字，它用来执行下一个case的body部分。这和C#，Java等语言不一样
	// Go默认只执行一条case中的body部分
	fmt.Println("There is cavern entrance here and a path to the east.")
	var command3 = "go inside"
	switch command3 {
	case "go east":
		fmt.Println("You head further up the mountain")
		fallthrough
	case "enter cave", "go inside":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case "read sign":
		fmt.Println("The sign reads 'No Minors'.")
	default:
		fmt.Println("Didn't quite get that.")
	}

	fmt.Println("==============")
	// for
	// for关键字可以让你的代码重复执行
	// 后边没有跟条件，那就是无限循环（可以使用break跳出循环）
	var count = 10
	for count > 0 {
		fmt.Println(count)
		// 停 1s
		time.Sleep(time.Second)
		count--
	}
	fmt.Println("Liftoff!")

}
