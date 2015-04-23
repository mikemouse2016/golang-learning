package main

import(
    "fmt"
    "time"
)
var c chan int
func ready(w string, sec int) {
time.Sleep(10)
fmt.Println(w, "is ready!")
c <- 1
}
func main() {
c = make(chan int)
go ready("Tee", 2)
go ready("Coffee", 1)
fmt.Println("I'm waiting, but not too long")
//<-c
//<-c
L: for {
    i := 0
    select {
case <-c:
    i++
    if i > 1 {
        break L
    }
    }
  }
}
/*
定义c 作为int 型的channel。就是说：这个channel 传输整数。注意这个变量是全局
的，这样goroutine 可以访问它；
.1. 发送整数1到channel c；
.2. 初始化c；
.3. 用保留字go 开始一个goroutine；
.4. 等待，直到从channel上接收一个值。注意，收到的值被丢弃了；
.5. 两个goroutines，接收两个值
虽然goroutine 是并发执行的，但是它们并不是并行运行的。如果不告诉Go 额外的东西，同
一时刻只会有一个goroutine 执行。利用runtime.GOMAXPROCS(n) 可以设置goroutine 并行执
行的数量
*/
