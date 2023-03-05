package main

import "fmt"

/**
知识点二：defer和return谁先谁后
答：return在前，defer在后
*/

func deferFunc() int {
	fmt.Println("defer func called...")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called...")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()

	return returnFunc()
}

func main() {
	returnAndDefer()
}
