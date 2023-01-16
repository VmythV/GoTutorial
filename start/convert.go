package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 64
	fmt.Println(x)                      // 64
	fmt.Println(string(x))              // @
	fmt.Println(strconv.Itoa(64))       // 64
	fmt.Println(strconv.Itoa(64) + "s") // 64s

	fmt.Println("--------------")

	// 另一种把数值转化为string的方式是使用Sprintf函数，和Printf略类似，但是会返回一个string
	countdown := 9
	str := fmt.Sprintf("Launch in T minus %v seconds.", countdown)
	fmt.Println(str)

	fmt.Println("--------------")

	// strconv包里面还有个Atoi（ASCII to Integer）函数
	// 由于字符串里面可能包含任意字符，或者要转化的数字字符串太大，所以Atoi函数可能会发生错误
	countdown2, err := strconv.Atoi("10")
	// 如果err的值为nil，那么就代表没发生错误
	if err != nil {
		// oh no, something went wrong
		fmt.Println(err)
		fmt.Println(err.Error())
	}
	fmt.Println(countdown2)

	// Go中无法将0，1 代表 false和true

}
