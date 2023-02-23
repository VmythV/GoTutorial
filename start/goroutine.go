package main

import (
	"fmt"
	"time"
)

// 在Go中，独立的任务叫做goroutine
// 虽然goroutine与其他语言中的协程，进程，线程都有相似之处，但goroutine和他们并不完全相同
// goroutine创建效率非常高
// Go能直截了当的系统多个并发（concurrent）操作
// 在某些语言中，将顺序式代码转化为并发式代码需要做大量修改
// 在Go里，无需修改现有顺序式代码，就可以通过goroutine以并发的方式运行任意数量的任务

// 每次使用go关键字都会产生一个新的goroutine
// 表面上看，goroutine似乎在同时运行，但由于计算机处理单元有限，其实技术上来说，这些goroutine不是真的在同时运行
// 计算机处理器会使用"分时"技术，在多个goroutine上轮流花费一些时间。
// 在使用goroutine时，各个goroutine的执行顺序无法确定

// 向goroutine传递参数就跟向函数传递参数一样，参数都是按值传递的（传入的是副本）

func main() {
	for i := 0; i < 5; i++ {
		go sleepyGopher(i) // 分支线路
	}
	time.Sleep(4 * time.Second) // 主线路
}

func sleepyGopher(id int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...snore...", id)
}
