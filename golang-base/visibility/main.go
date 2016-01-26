package main

import(
    "./v1"
    "./v2"
    "fmt"
)

func main(){
    v1.Demo()
//    v1.demo()  //error
    v2.Demo()
//    var d v2.demo  //error
    var test v2.Test
    test.B = 20
//    test.a = 10  //error
    fmt.Println(test)
}

/*
go语言中的符号访问性是包一级别的不是类型一级别的
同一个包中的型都可以访问同一包中的任意其他类型
*/
