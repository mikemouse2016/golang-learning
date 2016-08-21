package main

import (
	"fmt"
	"github.com/cosiner/zerver"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func HelloGet(req zerver.Request, resp zerver.Response) {
	u := &User{Id: 23, Name: "jianghuan"}
	//var a []byte
	//req.Read(u)
	//req.Read(a)
	fmt.Println(req.ContentType())
	fmt.Println(req.Param("id"))
	fmt.Println(req.RemoteAddr())
	fmt.Println(req.Pattern())
	fmt.Println(req.URL().Path)
	fmt.Println(req.Method())
	fmt.Println(req.AcceptEncodings())
	fmt.Println(req.Attrs())
	fmt.Println(req)
	fmt.Print(req.Attr("id"))
	//fmt.Println(a)
	resp.Send("user", u)
}

func HelloPost(req zerver.Request, resp zerver.Response) {
	u := &User{}
	req.Receive(u)
	fmt.Print(u)
	fmt.Println(req.ContentType())
	fmt.Println(req.Param("id"))
	fmt.Println(req.RemoteAddr())
	fmt.Println(req.Pattern())
	fmt.Println(req.URL().Path)
	fmt.Println(req.Method())
	fmt.Println(req.AcceptEncodings())
	fmt.Println(req.Attrs())
	fmt.Println(req)
	fmt.Print(req.Attr("id"))
	//resp.Send("user", u)
	resp.SetValue(u)
}

func main() {
	server := zerver.NewServer()
	server.Get("/", func(req zerver.Request, resp zerver.Response) {
		resp.WriteString("Hello World!")
	})
	server.Get("/hello", HelloGet)
	server.Post("/hello", HelloPost)
	server.Start(nil) // default listen at ":4000"
}
