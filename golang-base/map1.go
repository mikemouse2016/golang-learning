package main

import (
	"fmt"
)

type string_map map[string][]string

func main() {

	map1 := make(string_map)
	map1["123"] = []string{"abc","def"}
	
	s, ok := map1["321"]
	fmt.Println(ok)
	fmt.Println(s)
	
	if s, ok = map1["123"]; ok{
		fmt.Println(ok)
		fmt.Println(s)
	}
	map1["one"] = []string{"one", "1"} //先赋值
	one := map1["one"]                 //把赋值的映射使用一个变量表示
	one = []string{"two", "2"}         //给这个变量赋另外的值
	fmt.Println(one)
	fmt.Println(map1)

}
