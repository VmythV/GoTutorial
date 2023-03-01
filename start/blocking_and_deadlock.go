package main

// 当goroutine在等待通道的发送或接收时，我们就说它被阻塞了
// 除了goroutine本身占用少量的内存外，被阻塞的goroutine并不消耗任何其他资源
// goroutine静静地停在那里，等到导致其阻塞的事情来解除阻塞
// 当一个或多个goroutine因为某些永远无法发生的事情被阻塞时，我们称这种情况为死锁。而出现死锁的程序通常会崩溃或挂起

func main() {

}
