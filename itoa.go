package main

import (
	"fmt"
)

const (
	a = iota
	b = iota
	c = 30
	d = iota
)

func main() {
	fmt.Println(a, b, c, d)
	//s[0] = 'd'           直接修改字符串会报错
	s := "1你好"
	fmt.Println(s[1])
	fmt.Println(len(s))
	e := []rune(s)
	fmt.Println(len(e))
	fmt.Println(e[1])
	fmt.Println(s)
	e[1] = '我'
	fmt.Println(e[1])
	s1 := string(e)
	fmt.Println(s1)
}
