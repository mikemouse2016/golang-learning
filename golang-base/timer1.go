package main

import (
	"fmt"
	"time"
)

//使用time.NewTimer或time.NewTicker来设置一个定时器，这个定时器会绑定在当前channel中，通过channel的阻塞通知机器来通知程序
func main() {

	//Timer只通知一次
	timer := time.NewTimer(2 * time.Second)

	<-timer.C
	fmt.Println("timer expired!")
}
