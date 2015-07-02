package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)

	go func() {
		for t := range ticker.C {
			fmt.Println(t)
		}

	}()

	//设置一个timer，3钞后停掉ticker
	timer := time.NewTimer(3 * time.Second)
	<-timer.C

	ticker.Stop()

	fmt.Println("timer expired")
}
