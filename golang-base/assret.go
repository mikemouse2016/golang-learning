package main

import (
	"fmt"
)

type Html []interface{}

//断言的使用方法
func main() {
	html := make(Html, 6)
	html[0] = "div"
	html[1] = "span"
	html[2] = []byte("script")
	html[3] = "style"
	html[4] = "head"
	html[5] = 23
	for index, element := range html {
		if value, ok := element.(string); ok {
			fmt.Printf("html[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.([]byte); ok {
			fmt.Printf("html[%d] is a []byte and its value is %s\n", index, string(value))
		}

	}

	var val interface{} = "good"
	fmt.Println(val.(string))
	//fmt.Println(val.(int))
	//代码中被注释的那一行会运行时错误。这是因为val实际存储的是string类型，因此断言失败。

	//推荐使用
	for index, element := range html {
		switch value := element.(type) {
		case string:
			fmt.Printf("html[%d] is a string and its value is %s\n", index, value)
		case []byte:
			fmt.Printf("html[%d] is a []byte and its value is %s\n", index, string(value))
		case int:
			fmt.Printf("error type\n")
		default:
			fmt.Printf("unknown type\n")
		}
	}

}
