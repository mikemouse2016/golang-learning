package main

import "fmt"
import "time"

func main() {

	channel := make(chan string) //注意: buffer为1

	go func() {
		channel <- "hello"
		fmt.Println("write \"hello\" done!")

		channel <- "World" //Reader在Sleep，这里在阻塞
		fmt.Println("write \"World\" done!")

		fmt.Println("Write go sleep...")
		time.Sleep(3 * time.Second)
		channel <- "channel"
		fmt.Println("write \"channel\" done!")
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Reader Wake up...")

	msg, b := <-channel
	fmt.Println(b)
	fmt.Println("Reader: ", msg)

	msg, b = <-channel
	fmt.Println(b)
	fmt.Println("Reader: ", msg)

	msg, b = <-channel //Writer在Sleep，这里在阻塞
	fmt.Println(b)
	fmt.Println("Reader: ", msg)
}
