package main

import "fmt"

var c = make(chan int)   //无缓冲   
//无缓冲的不仅仅是向c通道放'c'而是 一直要有别的协程<-c接手了 这个参数，那么c<-'c'才会继续下去，要不然就一直阻塞着

func f() {
	c <- 'c'
	fmt.Println("在goroutine内")
}

func main() {
	go f()
	<-c
	fmt.Println("外部调用")
}
