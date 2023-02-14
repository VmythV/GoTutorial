package main

import "fmt"

// map是Go提供的另外一种集合
// 1 可以将key映射到value
// 2 快速通过key找到对应的value
// 3 key几乎可以是任意类型
// map[string]int    [string]:key type   int:value type

// 数组，int，float64等类型在赋值给新变量或传递至函数/方法的时候会创建相应的副本，但是map不会

// 除非你使用复合字面量来初始化map，否则必须使用内置的make函数来为map分配空间
// 创建map时，make函数可以接受一个或两个参数，第二个参数用于为指定数量的key预先分配空间

func main() {
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}

	temp := temperature["Earth"]
	fmt.Printf("On average the Earth is %v C.\n", temp)
	temperature["Earth"] = 16
	temperature["Venus"] = 464

	fmt.Println(temperature)

	moon := temperature["Moon"]
	fmt.Println(moon)

	fmt.Println("=================")

	planets := map[string]string{
		"Earth": "Sector ZZ9",
		"Mars":  "Sector ZZ9",
	}

	planetsMarkII := planets
	planets["Earth"] = "whoops"

	fmt.Println(planets)
	fmt.Println(planetsMarkII)

	// 删除map的中的元素
	delete(planets, "Earth")
	fmt.Println(planetsMarkII)
	fmt.Println(planets)

	fmt.Println("=================")
	//temperature2 := make(map[float64]int)
	temperature2 := make(map[float64]int, 8)

	fmt.Println(len(temperature2))

	fmt.Println("=================")

	temperature3 := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	frequency := make(map[float64]int)

	// 遍历赋值，如果key存在，则++，如果不存在，则创建并加一
	for _, t := range temperature3 {
		frequency[t]++
	}

	// 遍历时，无法保证顺序
	for n, num := range frequency {
		fmt.Printf("%+.2f occurs %d times\n", n, num)
	}

}
