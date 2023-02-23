package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 如何优雅的处理错误
// 减少错误处理代码的一种策略是：将程序中不会出错的部分和包含潜在错误隐患的部分隔离开来。
// 对于不得不返回错误的代码，应尽力简化相应的错误处理代码

// 使用defer关键字，Go可以确保所有deferred的动作可以在函数返回前执行
// 可以defer任意的函数和方法
// defer并不是专门做错误处理的
// defer可以消除必须时刻惦记执行资源释放的负担

func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return nil
	}
	// 延迟执行，保证方法执行完之前会调用
	defer f.Close()

	_, err = fmt.Println(f, "Errors are values.")
	if err != nil {
		return err
	}

	_, err = fmt.Println(f, "Don't just check errors, handle them gracefully")
	return err
}

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	fmt.Println("===============")
	err = proverbs("proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
