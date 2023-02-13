package main

import "fmt"

// append函数是内置函数，他可以将元素添加到slice里面

func main() {
	dwarfs := []string{"Ceres", "Pluto", "Hauma", "Makemake", "Eris"}

	dwarfs = append(dwarfs, "Orcus")

	dwarfs = append(dwarfs, "Salacia", "Quaoar", "Sedna")

	fmt.Println(dwarfs)

	fmt.Println("===================")

	dwarfs = []string{"Ceres", "Pluto", "Hauma", "Makemake", "Eris"}
	dump("dwarfs", dwarfs)
	dump("dwarfs[1:2]", dwarfs[1:2])

	fmt.Println("===================")
	planets := []string{"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune"}

	terrestrial := planets[0:4:4]
	worlds := append(terrestrial, "Ceres")

	dump("planets", planets)
	dump("terrestrial", terrestrial)
	dump("worlds", worlds)

	fmt.Println("===================")
	// 使用make函数，可以对slice进行预分配策略，尽量避免额外的内存分配和数组复制操作
	// make两个参数，第二个就是长度和容量；三个参数，第二个是长度，第三个是容量
	dwarfs2 := make([]string, 0, 10)
	dump("dwarfs2", dwarfs2)
}

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}
