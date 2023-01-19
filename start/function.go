package main

// func  Intn(   n      int)     int
// 关键字 函数名称 参数变量 参数类型 返回值类型

// num := rand.Intn(10)
//       包名  方法名 参数

// 在Go里，大写字母开头的函数，变量或其他标识符都会被导出，对其他包可用，小写字母开头的就不行

// func myFunc(n int ,m int) int

// func myFunc(n,m int) (i int,err error)
// func myFunc(n,m int) (int,error)
// 函数的多个返回值需要用括号括起来，每个返回值名字在前，类型在后。声明函数时可以把名字去掉，只保留类型

// 可变参数函数，Println：func Println(a ...interface{}) (n int,err error)    interface{} 是一个空接口，可以接受任意类型的参数

func myFunc1(n int) int {
	return n
}

func myFunc2(n int, m int) int {
	return n + m
}

func myFunc3(n int) (i int, j int) {
	return n, n * n
}
func myFunc4(n int) (int, int) {
	return n, n * n
}

func myFunc5(n ...interface{}) (int, int) {
	return 1, 2
}

func main() {

}
