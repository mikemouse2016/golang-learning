package main

import(
    "fmt"
)


func test(){
    fmt.Println(222)
    return 
    defer func(){
        fmt.Println(333)
    }()
}

func main(){
    test()
}
