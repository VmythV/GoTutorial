package main

import "fmt"

// 数组是一种固定长度的且有序的元素集合。
// var planets [8]string
// Go编译器在检测到对越界元素的访问时会报错
// 如果Go编译器在编译时未能发现越界错误，那么程序在运行时会出现panic，这时程序就会崩溃
func main() {
	var planets [8]string

	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"

	earth := planets[2]
	fmt.Println(earth)

	fmt.Println(len(planets))
	fmt.Println(planets[3] == "")

	// =================================

	planets2 := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	fmt.Println(planets2)

	// =================================

	dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	for i := 0; i < len(dwarfs); i++ {
		fmt.Println(i, dwarfs[i])
	}

	fmt.Println("========")

	for i, dwarf := range dwarfs {
		fmt.Println(i, dwarf)
	}

	// =================================

	// 类似值传递
	fmt.Println("========")
	planets3 := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	terraform(planets3)
	fmt.Println(planets3)

}

func terraform(planets3 [8]string) {
	for i := range planets3 {
		planets3[i] = "New " + planets3[i]
	}
	fmt.Println(planets3)
}
