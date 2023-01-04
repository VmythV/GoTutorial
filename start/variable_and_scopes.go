package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 变量声明的位置，决定了它的作用域
// 作用域的好处：可以在不同的作用域内使用相同的变量名
// 在Go中，作用域的范围就是{}之间的部分

// 在这边声明的变量，拥有package作用域，如果main package 有多个函数，era对他们都可见
// 注意：短声明不可用来声明package作用域的变量（ a:=10 这种就是短声明）
var era = "AD"

func main() {
	// count的作用域就是在main里，num的作用域在for里
	var count = 0
	for count < 10 {
		var num = rand.Intn(10) + 1
		fmt.Println(num)

		count++
	}

	fmt.Println("====================")

	// 使用短声明
	for count := 10; count > 0; count-- {
		fmt.Println(count)
	}

	fmt.Println("====================")
	// 在if中使用短声明
	rand.Seed(time.Now().Unix())
	if num := rand.Intn(3); num == 0 {
		fmt.Println("Space Adventures")
	} else if num == 1 {
		fmt.Println("SpaceX")
	} else {
		fmt.Println("Virgin Galactic")
	}

	fmt.Println("====================")
	// 在switch中使用短声明
	switch num := rand.Intn(10); num {
	case 0:
		fmt.Println("Space Adventures")
	case 1:
		fmt.Println("SpaceX")
	case 2:
		fmt.Println("Virgin Galactic")
	default:
		fmt.Println("Random spaceline #", num)
	}

	fmt.Println("====================")
	year := 2018

	switch month := rand.Intn(12) + 1; month {
	case 2:
		day := rand.Intn(28) + 1
		fmt.Println(era, year, month, day)
	case 4, 6, 9, 11:
		day := rand.Intn(30) + 1
		fmt.Println(era, year, month, day)
	default:
		day := rand.Intn(31) + 1
		fmt.Println(era, year, month, day)
	}
}
