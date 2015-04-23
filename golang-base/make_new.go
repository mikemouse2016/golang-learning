package main

import(
    "fmt"
)

func  main(){
    var p *[]int = new([]int)
    var v []int = make([]int, 10, 100)
    *p = make([]int, 5, 10)
    fmt.Println(p, v, len(v), cap(v),*p, len(*p), cap(*p))
    
}
/*
make 仅适用于map，slice 和channel，并且返回的不是指针。应当用new 获得特定
的指针。
new(T) 返回*T 指向一个零值T
make(T) 返回初始化后的T
*/
