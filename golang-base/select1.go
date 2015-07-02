package main

import (
	"fmt"
	"time"
)

func main() {

	//创建两个channel
	c1 := make(chan string)
	c2 := make(chan string)

	//创建两个goruntime分别向两个channel发送数据
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "Hello"
	}()

	go func() {
		time.Sleep(time.Second * 1)
		c2 <- "World"
	}()

	/*
		//用select来侦听两个channel
		for i := 0; i < 2; i++ {
			select {
			case msg1 := <-c1:
				fmt.Println("received", msg1)
			case msg2 := <-c2:
				fmt.Println("received", msg2)
			}
		}
	*/
	/*
		//通过timeout
		timeout_cnt := 0
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("msg1 received", msg1)
			case msg2 := <-c2:
				fmt.Println("msg2 received", msg2)
			case <-time.After(time.Second * 3):
				fmt.Println("Time Out")
				timeout_cnt++
			}

			if timeout_cnt > 3 {
				break
			}
		}
	*/

	for {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		default:
			fmt.Println("nothing received!")
			time.Sleep(time.Second)
		}
	}

	fmt.Println("over")
}
