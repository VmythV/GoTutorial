package main

import "fmt"

// 切片的扩容机制：append的时候，如果长度增加后超过容量，则将容量增加2倍

func main() {
	var numbers = make([]int, 3, 5)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 向numbers切片追加一个元素1，numbers len = 4，[0,0,0,1] cap = 5
	numbers = append(numbers, 1)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 向numbers切片追加一个元素2，numbers len = 5，[0,0,0,1] cap = 5
	numbers = append(numbers, 1)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 向一个容量cap已经满的slice追加元素
	numbers = append(numbers, 1)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	fmt.Println("============")

	var numbers2 = make([]int, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)
	numbers2 = append(numbers2, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)

}
