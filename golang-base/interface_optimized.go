package main

import(
    "fmt"
    "time"
)
type Tester interface {
    Test(int)
}

type Data struct {
    x int
}

func (d *Data) Test(x int) {
    d.x += x
}

func call(d *Data) {
    d.Test(100)
}

func interfaceCall(t Tester){
    t.Test(100)
}

func main(){
    d := &Data{x:100}
    t1 := time.Now().UnixNano()
    for i := 0; i<10000000; i++{
        call(d)
    }
    t2 := time.Now().UnixNano()
    fmt.Println(t2-t1)
    for i := 0; i<10000000; i++{
        interfaceCall(d)
    }
    t3 := time.Now().UnixNano()
    fmt.Println(t3-t2)
    fmt.Println("%+v", d)
}    
