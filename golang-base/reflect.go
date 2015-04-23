package main

import (
    "fmt"
    "reflect"
)

type Person struct {
    name string "namestr" //  "namestr" 是标签
    age int
}

/*
ShowTag(p1)   //调用ShowTag() 并传递指针
func ShowTag(i interface{}) {
    switch t := reflect.TypeOf(i); t.Kind() {    //Get type, switch on Kind()
    case reflect.Ptr:   //Its a pointer, hence a reflect.Ptr
        tag := t.Elem().Field(0).Tag
*/

func show(i interface{}) {

    switch t := i.(type) {
    case *Person:
        tt := reflect.TypeOf(i)   //得到类型的元数据
        v := reflect.ValueOf(i)   //得到实际的值
        tag := tt.Elem().Field(0).Tag
        name := v.Elem().Field(0).String()
        fmt.Println(tt, v, tag, name, t)
    }
}

func main() {
    p1 := new(Person)   //new 返回Person 的指针
    p1.name = "jianghuan"
    p1.age = 10
    show(p1)
    fmt.Println(p1)
}
