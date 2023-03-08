package main

import "fmt"

func printArray2(myArray []int) {
	// 引用传递，可以改变原始对象的值
	// _ 表示匿名的变量
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
	myArray[0] = 100
}

func main() {
	// 动态数组，切片 slice
	myArray := []int{1, 2, 3, 4}

	fmt.Printf("myArray type is %T\n", myArray)

	printArray2(myArray)

	fmt.Println(" ======== ")

	for _, value := range myArray {
		fmt.Println("value = ", value)
	}

}
