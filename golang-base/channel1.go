package main

import (
	"fmt"
)

func main() {
	//创建string类型的channel
	channel := make(chan string)

	//创建一个goruntine想channel发一个字符串

	go func() {
		fmt.Println(1)
		channel <- "hello"
	}()

	fmt.Println(2)
	msg := <-channel
	fmt.Println(3)

	fmt.Println(msg)
}

/* channel默认是阻塞的，也就是channel满了就默认阻塞写，channel空了，就默认阻塞读
   可利用此特性同步数据
*/
