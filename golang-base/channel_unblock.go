package main

import "fmt"

var c = make(chan int, 1) //有缓冲
//c <- 'c'则不会阻塞，因为缓冲大小是1 只有当放第二个值的时候 第一个还没被拿走，这时候才会阻塞。

func f() {
	c <- 'c'
	fmt.Println("在goroutine内")
}

func main() {
	go f()
	c <- 'c'
	<-c
	<-c
	fmt.Println("外部调用")
}
