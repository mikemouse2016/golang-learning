package main

import (
	"fmt"
)

func main() {
	a := []int{}
	fmt.Println(len(a), cap(a), a)
	a = append(a, 1)
	fmt.Println(len(a), cap(a), a)
	a = append(a, 10)
	fmt.Println(len(a), cap(a), a)
}
