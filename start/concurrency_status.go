package main

import "sync"

var mu sync.Mutex

func main() {
	mu.Lock()
	// 互斥锁
	defer mu.Unlock()
}
