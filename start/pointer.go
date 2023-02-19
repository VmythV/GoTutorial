package main

import "fmt"

// 指针是指向另一个变量地址的变量
// Go语言的指针同时也强调安全性，不会出现迷途指针（dangling pointers）

// 变量会将它们的值存储在计算机的RAM里，存储位置就是该变量的内存地址
// &表示地址操作符，通过&可以获得变量的内存地址
// &操作符无法获取字符串/数值/布尔字面值的地址（&42，&"hello"这些都会导致编译器报错 ）
// *操作符与&的作用相反，它用来解引用，提供内存地址指向的值

// 指针存储的是内存地址
// 指针类型和其他普通类型一样，出现在所有需要用到类型的地方，如变量声明、函数形参、返回值类型、返回值类型、结构体字段等

// 将*放在类型前面表示声明指针类型
// 将*放在变量前面表示解引用操作

// 两个指针变量持有相同的内存地址，那么它们就是相等的。

// 与字符串和数值不一样，复合字面量的前面可以放置&

// 访问字段时，对结构体进行解引用不是必须的

// =====
// 和结构体一样，可以把&放在数组的复合字面量值前面来创建指向数组的指针
// 数组在执行索引或切片操作时会自动解引用，没有必要写(*superpower)[0]这种形式
// 与C语言不一样，Go里面数组和指针是两种完全独立的类型
// Slice和map的复合字面值前面也可以放置&操作符，但Go并没有为它们提供自动解引用的功能
func main() {
	answer := 42
	fmt.Println(&answer)

	address := &answer
	fmt.Println(*address)
	// %T 打印变量的地址
	fmt.Printf("address is a %T\n", address)

	fmt.Println("============")

	type person struct {
		name, superpower string
		age              int
	}

	timmy := &person{
		name: "Timothy",
		age:  10,
	}

	// 效果是一样的
	(*timmy).superpower = "flying"
	timmy.superpower = "flying"

	fmt.Printf("%+v\n", timmy)

	fmt.Println("============")
	superpower := &[3]string{"flight", "invisibility", "super strength"}

	fmt.Println(superpower[0])
	fmt.Println(superpower[1:2])
}
