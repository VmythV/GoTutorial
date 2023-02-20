package main

import (
	"fmt"
	"strings"
)

// Go语言的函数和方法都是按值传递参数的，这意味着函数总是操作于被传递参数的副本
// 当指针被传递到函数时，函数将接收传入的内存地址的副本。之后函数可以通过解应用内存地址来修改指针指向的值

// 方法的接收者和方法的参数在处理指针方面是很相似的
// Go语言在变量通过点标记法进行调用的时候，自动使用&取变量的内存地址

// Go语言提供了 内部指针 这种特性
// 它用于确定结构体中指定字段的内存地址
// &操作符不仅可以获取结构体的内存地址，还可以获得结构体中指定字段的内存地址

// 函数通过指针对数组的元素进行修改

// Go语言里一些内置的集合类型就在暗中使用指针
// map在被赋值或者被作为参数传递的时候不会被复制（map就是一种隐式指针）
// map的键和值都可以是指针类型
// 需要将指针指向map的情况并不多见

// slice是指向数组的窗口，实际上slice在指向数组元素的时候也使用了指针
// 每个slice内部都会被表示为一个包含3个元素的结构，它们分别指向 数组的指针，slice的容量，slice的长度
// 当slice被直接传递至函数或方法时，slice的内部指针就可以对底层数据进行修改
// 指向slice的显式指针的唯一作用就是修改slice本身：slice的长度，容量以及起始偏移量

type person struct {
	name, superpower string
	age              int
}

func birthday(p *person) {
	p.age++
}

func (p *person) birthday() {
	p.age++
}

type stats struct {
	level             int
	endurance, health int
}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 + s.level)
	s.health = 5 * s.endurance
}

type character struct {
	name  string
	stats stats
}

func reset(board *[8][8]rune) {
	board[0][0] = 'r'
}

func reclassify(planets *[]string) {
	*planets = (*planets)[0:8]
}

type talker2 interface {
	talk() string
}

func shout2(t talker2) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type martian2 struct{}

func (m martian2) talk() string {
	return "nack nack"
}

type laser2 int

func (l *laser2) talk() string {
	return strings.Repeat("pew ", int(*l))
}

func main() {
	rebecca := person{
		name:       "Rebecca",
		superpower: "imagination",
		age:        14,
	}

	birthday(&rebecca)

	fmt.Printf("%+v\n", rebecca)

	//(&rebecca).birthday()
	rebecca.birthday()

	fmt.Printf("%+v\n", rebecca)

	fmt.Println("===================")

	player := character{name: "matthias"}
	levelUp(&player.stats)

	fmt.Printf("%+v\n", player.stats)

	fmt.Println("===================")
	var board [8][8]rune
	reset(&board)
	fmt.Printf("%c", board[0][0])

	fmt.Println("===================")
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
		"Pluto",
	}
	reclassify(&planets)

	fmt.Println(planets)

	fmt.Println("===================")
	// 效果一样
	shout2(martian2{})
	shout2(&martian2{})

	fmt.Println("===================")
	pew := laser2(2)
	shout2(&pew)
}
