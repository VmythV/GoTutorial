package main

import "fmt"

func main() {
	// Go中字符串中可以包含转义字符，但如果就想要打印转义字符，那就用 `` 来包含字符串
	fmt.Println("pace be upon you\nupon you be peace")
	fmt.Println(`string can span multiple lines with the \n escape sequence`)
	fmt.Println(`
			pace be upon you
			upon you be peace`)

	fmt.Println("----------")

	// 别名就是同一个类型的另一个名字
	// 为了表示这样的unicode code point，Go语言提供rune这个类型，他是int32的一个类型别名
	// byte是uint8 类型的别名，目的是用于二进制数据

	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33

	// 打印出 code point 的值
	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)

	// 打印出字符
	fmt.Printf("%c%c%c%c\n", pi, alpha, omega, bang)

	// string是不可变的
	// 加密算法，将每个字母都向后移动几位，例如：a ——> d，b ——> e

	c := 'a'
	c = c + 3
	fmt.Printf("%c", c)
	if c > 'z' {
		c = c - 26
	}

	fmt.Println("----------")
	// Go中有很多内置函数，len("xxx") 就是返回所占的byte数，内置函数不需要import
	// Go里的函数可以返回多个值
	ROT13()
}

func ROT13() {
	message := "uv vagreangvbany fcnpr fgngvba"

	//for i := 0; i < len(message); i++ {
	//	c := message[i]
	//	if c >= 'a' && c <= 'z' {
	//		c = c + 13
	//		if c > 'z' {
	//			c = c - 26
	//		}
	//	}
	//	fmt.Printf("%c", c)
	//}

	// range可以遍历各种集合
	for i := range message {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}

	//for i, c := range "hello" {
	//	fmt.Printf("%v %c\n", i, c)
	//}
	// 也可以丢弃到某些值
	//for _, c := range "hello" {
	//	fmt.Printf("%v \n", i)
	//}
}
