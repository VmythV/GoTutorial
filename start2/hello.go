package main // 程序的包名
import "fmt"

/*
	四种变量的声明方式
*/

// 声明全局变量 方法一，方法二，方法三是可以的，但是方法四不行
var gA int = 100
var gB = 200

func main() {
	// 方式一：声明一个变量 默认值是0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	// 方式二：声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	// 方式三：在初始化的时候，可以省去数据类型，通过值自动匹配当前的变量的数据类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	// 方式四：（常用方法）省去var关键字，直接自动匹配（只能在方法体内使用）
	d := 100
	fmt.Println("d = ", d)
	fmt.Printf("type of d = %T\n", d)

	// =====
	fmt.Println("gA = ", gA, "gB = ", gB)

	// 声明多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx = ", xx, "yy = ", yy)
	var kk, ll = 100, "ABC"
	fmt.Println("kk = ", kk, "ll = ", ll)

	// 多行的多变量声明c
	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv = ", vv, "jj = ", jj)
}
