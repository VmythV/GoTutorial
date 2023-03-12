package main

import "fmt"

func main() {
	cityMap := make(map[string]string)

	// 添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	// 遍历
	printMap(cityMap)

	// 删除
	delete(cityMap, "China")

	// 修改
	cityMap["USA"] = "DC"
	changeValue(cityMap)

	fmt.Println("-----------")

	// 遍历
	printMap(cityMap)
}

func changeValue(cityMap map[string]string) {
	cityMap["England"] = "London"
}

func printMap(cityMap map[string]string) {
	// cityMap 是一个引用传递
	for key, value := range cityMap {
		fmt.Println(key)
		fmt.Println(value)
	}
}
