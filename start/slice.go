package main

import (
	"fmt"
	"strings"
)

// 指向数组的窗口
// 假设planets是一个数组，那么planets[0:4]就是一个切片，他切分出了数组里前4个元素
// Go中索引不能是负数
// 数组可以切片，字符串也可以切片
// 数组按值传递，slice按引用传递

// Slice使用的是半开区间(左闭右开)
// 例如planets[0:4]，包含索引0，1，2，3对应的元素，不包括索引4对应的元素

func main1() {
	planets5 := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	// [:4]   [6:]   [:]
	terrestrial := planets5[0:4]
	gasGiants := planets5[4:6]
	iceGiants := planets5[6:8]

	fmt.Println(terrestrial, gasGiants, iceGiants)

	fmt.Println("===================")

	// 改变原来切片的值不会改变切片后的值
	neptune := "Neptune"
	tune := neptune[3:]

	fmt.Println(tune)

	neptune = "Poseidon"
	fmt.Println(tune)

	fmt.Println("===================")

	dwarfArray := [...]string{"Ceres", "Pluto", "haumea", "Makemake", "Eris"}

	dwarfSlice := dwarfArray[:]

	dwarfs := []string{"Ceres", "Pluto", "haumea", "Makemake", "Eris"}

	fmt.Println(dwarfSlice, dwarfs)

}

func main() {
	planets := []string{"Venus    ", "Earth  ", " Mars"}
	hyperspace(planets)

	fmt.Println(strings.Join(planets, "-"))
}

func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}
