package main

import (
	"fmt"
	"runtime"
)

func f(msg string) {
	fmt.Println(msg)
}

func main() {

	runtime.GOMAXPROCS(4)
	go f("go runtime")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
