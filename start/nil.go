package main

import (
	"fmt"
	"sort"
)

// Nil是一个名称，表示"无"或者"零"。在Go中，nil是一个零值
// 如果一个指针没有明确对的指向，那么它的值就是nil
// 除了指针，nil还是slice、map和接口的零值

type person2 struct {
	age int
}

func (p *person2) birthday() {
	if p == nil {
		return
	}
	p.age++
}

func sortStrings(s []string, less func(i, j int) bool) {
	if less == nil {
		less = func(i, j int) bool { return s[i] < s[j] }
	}
	sort.Slice(s, less)
}

func main() {
	var nobody *person2
	fmt.Println(nobody)

	nobody.birthday()

	fmt.Println("==========")
	food := []string{"onion", "carrot", "celery"}
	sortStrings(food, nil)
	fmt.Println(food)
}
