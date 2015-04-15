package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {

	// 解析参数，默认是不会解析的
	//r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("URL", r.URL)
	fmt.Println("Proto", r.Proto)
	fmt.Println("Method", r.Method)
	fmt.Println("ContentLength", r.ContentLength)
	fmt.Println("RemoteAddr", r.RemoteAddr)
	fmt.Println("Host", r.Host)
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello world")
}

func main() {

	// 设置访问的路由
	http.HandleFunc("/", sayHelloName)

	// 设置监听的端口
	err := http.ListenAndServe(":9000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
