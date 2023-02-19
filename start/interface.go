package main

import (
	"fmt"
	"strings"
)

// 接口关注于类型可以做什么，而不是存储了什么
// 接口通过列举类型必须满足的一组方法来进行声明
// 在Go语言中，不需要显示声明接口

var t interface {
	talk() string
}

type talker interface {
	talk() string
}

type martian struct{}

// 该方法实现了t的接口
func (m martian) talk() string {
	return "nack nack"
}

type laser int

// 该方法实现了t的接口
func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func main() {
	t = martian{}
	fmt.Println(t.talk())

	t = laser(3)
	fmt.Println(t.talk())

	shout(martian{})
	shout(laser(2))
}
