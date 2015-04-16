package main

import(
    "fmt"
)
type I interface {
Get() int
Put(int)
}
type S struct { i int }
func (p *S) Get() int { return p.i }
func (p *S) Put(v int) { p.i = v }

type R struct { i int }
func (p *R) Get() int { return p.i }
func (p *R) Put(v int) { p.i = v }

func f(p I) {
switch t := p.(type) {  //类型判断。在switch 语句中使用(type)。保存类型到变量t,在switch 之外使用(type) 是非法的
case *S:
    fmt.Println("*S")
case *R:
    fmt.Println("*R")
//case S:
//    fmt.Println("S")
//case R:
//    fmt.Println("R")
default:
    fmt.Println("default", t)
}
}

func main(){
    var s S; f(&s)
}

/*
确定一个变量实现了某个接口，可以使用：
t := any.(I)
*/
