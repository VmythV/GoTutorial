package main

import (
	"fmt"
	"time"
)

// 通道（channel）可以在对个goroutine之间安全的传值。
// 通道可以用作变量，函数参数，结构体字段...
// 创建通道用make函数，并指定其传输数据的类型（c := make(chan int)）

// 通道channel发送，接受
// 使用左箭头操作符 <- 向通道发送值 或 从通道接收值
// 向通道发送值： c <- 99
// 从通道接收值： r := <- c
// 发送操作会等待直到另一个goroutine尝试对该通道进行接收操作为止
// 执行发送操作的goroutine在等待期间将无法执行其他操作
// 未在等待通道操作的goroutine任然可以继续自由的运行
// 执行接收操作的goroutine将等待直到另一个goroutine尝试向该通道进行发送操作为止。

// 使用select处理多个通道
// 等待不同类型的值
// time.After函数，返回一个通道，该通道在指定时间后会接收到一个值（发送该值的goroutine是Go运行时的一部分）
// select和switch有点像
// 该语句包含的每个case都持有一个通道，用来发送或接收数据
// select会等待直到某个case分支的操作者就绪，然后就会执行该case分支
// 注意：即使已经停止等待goroutine，但只要main函数还没返回，任在运行的goroutine将会继续占用内存
// select语句在不包含任何case的情况下将永远等下去

// nil通道
// 如果不使用make初始化通道，那么通道变量的值就是nil（零值）
// 对nil通道进行发送或者接收不会引起panic，但会导致永久阻塞。
// 对nil通道执行close函数，那么会引起panic
// nil通道的用处：
// 对于包含select语句的循环，如果不希望每次循环都等待select所涉及的所有通道，那么可以先将某些通道设为nil，等到发送值准备就绪之后，再将通道变成一个非nil值并执行发送操作。

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGopher2(i, c)
	}
	for i := 0; i < 5; i++ {
		gopherID := <-c
		fmt.Println("gopher ", gopherID, " has finished sleeping")
	}
}

func sleepyGopher2(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("... ", id, " snore")
	c <- id
}
