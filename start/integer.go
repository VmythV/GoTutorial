package main

import (
	"fmt"
	"math"
)

// Go中提供了10种整数类型：不可以存小数，范围有限，通常根据数值范围来选取整数类型
// 5种整数类型是有符号的，能表示正数，0，负数
// 5中整数类型是无符号的，能表示正数，0

// 整数类型取值范围各不相同，但与架构无关
// int8    -128 to 127   											8位
// uint8   0 to 255      											8位
// int16   -32,768 to 32,767   										16位
// uint16  0 to 65,535      										16位
// int32   -2,147,483,648 to 2,147,483,647   						32位
// uint32  0 to 4,294,967,295      									32位
// int64   -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807	64位
// uint64  0 to 18,446,744,073,709,551,615      					64位

// int和uint是针对目标设备优化的类型：
// 在树莓派2，比较老的移动设备上，int和uint都是32位的
// 在比较新的计算机上，int和uint都是64位的。
func main() {
	// 有符号的整数类型
	var year int = 2018
	// 无符号的整数类型
	var month uint = 2

	// 使用 %T 可以显示该变量的类型
	fmt.Printf("Type %T for %v\n", year, year)
	fmt.Printf("Type %T for %v\n", month, month)

	fmt.Println("---------------------------------")

	// 建议颜色的数据用uint8来表示
	var red1, green1, blue1 uint8 = 0, 141, 213
	// 在数前面加上0x前缀，可以用十六进制的形式表示数值
	var red2, green2, blue2 uint8 = 0x00, 0x8d, 0xd5
	fmt.Printf("%x %x %x\n", red1, green1, blue1)
	fmt.Printf("color: #%02x%02x%02x\n", red2, green2, blue2)

	fmt.Println("---------------------------------")

	// 所有整数类型都有一个取值范围，超出这个范围，就会发生环绕
	var red uint8 = 255
	red++
	fmt.Println(red)

	var number int8 = 127
	number++
	fmt.Println(number)

	fmt.Println("---------------------------------")

	// 打印数字的bit，使用%b
	var green uint8 = 3
	fmt.Printf("%08b\n", green)
	green++
	fmt.Printf("%08b\n", green)

	fmt.Println("---------------------------------")

	// 整数类型的最大最小值
	fmt.Println(math.MaxInt16)
}
