package main

import (
	"fmt"
	"reflect"
)

type Info struct {
	name string `abc:"type,attr,omitempty" nnn:"xxx"`
	pass int    `test`
}

func main() {
	info := Info{"hello", 1}
	ref := reflect.ValueOf(info)
	fmt.Println(ref.Kind())
	fmt.Println(reflect.Interface)
	fmt.Println(ref.Type())
	typ := reflect.TypeOf(info)
	n := typ.NumField()
	fmt.Println(n)
	for i := 0; i < n; i++ {
		f := typ.Field(i)
		fmt.Println(f.Tag)
		fmt.Println(f.Tag.Get("nnn"))
		fmt.Println(f.Name)
	}
}
