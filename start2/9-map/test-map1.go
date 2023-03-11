package main

import "fmt"

func main() {
	// 第一种声明方式
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1 是一个空map")
	}

	// 在使用map前，需要先用make给map分配数据空间
	myMap1 = make(map[string]string, 10)

	myMap1["one"] = "java"
	myMap1["two"] = "c++"
	myMap1["three"] = "python"

	fmt.Println(myMap1)

	// 第二种声明方式
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "python"

	fmt.Println(myMap2)

	// 第三种声明方式
	myMap3 := map[string]string{
		"one":   "java",
		"two":   "c++",
		"three": "python",
	}

	fmt.Println(myMap3)

}
