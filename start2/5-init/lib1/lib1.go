package lib1

import "fmt"

func init() {
	fmt.Println("lib1 init() ... ")
}

// 当前包提供的API
func Lib1Test() {
	fmt.Println("lib1Test() ... ")
}
