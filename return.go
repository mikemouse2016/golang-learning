package main

import (
	"fmt"
)

/*
Go 函数的返回值或者结果参数可以指定一个名字,并且像原始的变量那样使用,就像
输入参数那样。如果对其命名,在函数开始时,它们会用其类型的零值初始化。如果
函数在不加参数的情况下执行了 return 语句,结果参数会返回。
*/
func ReadFull() (n int, err error) {
	err = nil
	n = 3
	fmt.Println(err, n)
	return n, err
}

func main() {
	ReadFull()
}
