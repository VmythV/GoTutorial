package main

// 使用rand包，可以生成伪随机数
// 例如，Intn可以返回一个指定范围的随机整数
// import 的路径是 "math/rand"

//import "fmt"
//import "math/rand"

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)
}
