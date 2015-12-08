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
/*
Go中的数组是数值，因此当你向函数中传递数组时，函数会得到原始数组数据的一份复制
更新原始数组的数据，可以使用数组指针类型
*/
        x := [3]int{1,2,3}
        func(arr [3]int) {
            arr[0] = 7
            fmt.Println(arr) //prints [7 2 3]
        }(x)
        fmt.Println(x) //prints [1 2 3] (not ok if you need [7 2 3])

        func(arr *[3]int) {
            (*arr)[0] = 7
            fmt.Println(arr) //prints &[7 2 3]
        }(&x)
        fmt.Println(x) //prints [7 2 3]
}

