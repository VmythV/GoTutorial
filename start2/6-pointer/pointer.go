package main

import (
	"fmt"
)

//func swap(a int, b int) {
//	var temp int
//	temp = a
//	a = b
//	b = temp
//}

func swap(pa *int, pb *int) {
	var temp int
	temp = *pa
	*pa = *pb
	*pb = temp
}

func main() {
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Println("a = ", a, "b = ", b)

	var p *int
	p = &a
	fmt.Println(&a) // 二级指针
	fmt.Println(p)

	var pp **int
	pp = &p
	fmt.Println(&p)
	fmt.Println(pp)
}
