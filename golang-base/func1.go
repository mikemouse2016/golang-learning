package main

import (
    "fmt"
)

func demo(cmd func(int)int){
fmt.Println(cmd(23))
fmt.Println(cmd)
}

func demo1(i int)int{
    return i
}

func main(){
    demo(demo1)
}
