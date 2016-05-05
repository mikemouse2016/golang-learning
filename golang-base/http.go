package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName1(w http.ResponseWriter, r *http.Request) {

	// 解析参数，默认是不会解析的
	r.ParseForm()
        fmt.Println(1111)
	fmt.Println(r.Form)
	fmt.Println("URL", r.URL)
	fmt.Println("Proto", r.Proto)
	fmt.Println("Method", r.Method)
	fmt.Println("ContentLength", r.ContentLength)
	fmt.Println("RemoteAddr", r.RemoteAddr)
	fmt.Println("Host", r.Host)
	fmt.Println(r.Header.Get("X-Real-IP"))
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello world")
}

func sayHelloName2(w http.ResponseWriter, r *http.Request) {

	// 解析参数，默认是不会解析的
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("URL", r.URL)
	fmt.Println("Proto", r.Proto)
	fmt.Println("Method", r.Method)
	fmt.Println("ContentLength", r.ContentLength)
	fmt.Println("RemoteAddr", r.RemoteAddr)
	fmt.Println("Host", r.Host)
	fmt.Println(r.Header.Get("X-Real-IP"))
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello world")
}

func sayHelloName3(w http.ResponseWriter, r *http.Request) {

	// 解析参数，默认是不会解析的
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("URL", r.URL)
	fmt.Println("Proto", r.Proto)
	fmt.Println("Method", r.Method)
	fmt.Println("ContentLength", r.ContentLength)
	fmt.Println("RemoteAddr", r.RemoteAddr)
	fmt.Println("Host", r.Host)
	fmt.Println(r.Header.Get("X-Real-IP"))
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello world")
}

func main() {

	// 设置访问的路由
	http.HandleFunc("/test1/", sayHelloName1)
	http.HandleFunc("/test2", sayHelloName2)
	http.HandleFunc("/friend/r/", sayHelloName3)

	// 设置监听的端口
	err := http.ListenAndServe("0.0.0.0:8000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
