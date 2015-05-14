package main

import(
	"fmt"
)
func type_switch(a interface{}){
	//第一种
	if inst1,ok := a.(myStruct); ok{
		fmt.Println(ok)
	    fmt.Println(inst1)
	}
	//第二种
	switch inst2 :=a.(type){
	        case myStruct:
	                fmt.Println(inst2)
	        default:
	                fmt.Println("unknow")
	
	}
}

type myStruct struct {
	text     string
	moreData []byte
}
/*
a可能是任意类型
a.(某个类型) 返回两个值inst和ok ，ok代表是否是这个类型，如果是Ok inst就是转换后的类型 
a.(type) type是关键字 结合switch case使用
TypeA(a) 是强制转换
*/
func main(){
	m := myStruct{"1", []byte{1,2,3}}
	type_switch(m)
}
