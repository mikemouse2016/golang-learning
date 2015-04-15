package main

import (
	"fmt"
)

/*
Go 函数的返回值或者结果参数可以指定一个名字,并且像原始的变量那样使用,就像
输入参数那样。如果对其命名,在函数开始时,它们会用其类型的零值初始化。如果
函数在不加参数的情况下执行了 return 语句,结果参数会返回。
*/
func myfunc(arg ...int) {
	for _, n := range arg {
		fmt.Printf("And the number is: %d\n", n)
	}
	myfunc2(arg[:2]...)
}

func myfunc2(arg ...int) {
	for _, n := range arg {
		fmt.Printf("number is: %d\n", n)
	}
}

func main() {
	myfunc(1, 2, 3, 4)
	a := func() {
		println("Hello")
	}
	a()
	arxs := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		3: func() int { return 30 },
	}
	fmt.Println(arxs)
	callback(10, printit)
}

func callback(y int, f func(int)) {
	f(y)
}
func printit(x int) {
	fmt.Printf("%v\n", x)
}
