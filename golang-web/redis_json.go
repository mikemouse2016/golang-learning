package main

import (
	"fmt"
	"github.com/cosiner/zerver"
	"github.com/garyburd/redigo/redis"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func listmsg(req zerver.Request, resp zerver.Response) {
	conn, err := redis.DialTimeout("tcp", ":6379", 0, 1*time.Second, 1*time.Second)
	if err != nil {
		fmt.Println("error")
		return
	}
	defer conn.Close()
	fmt.Println(req.Param("id"))
	_, err = conn.Do("SET", "user", 123)
	_, err = conn.Do("SET", "user", 456)
	_, err = conn.Do("APPEND", "user:user0", 87)

	user0, err := redis.Int(conn.Do("GET", "user:user0"))
	user1, err := redis.Int(conn.Do("GET", "user:user1"))

	fmt.Printf("user0 is %d , user1 is %d \n", user0, user1)

	size, _ = conn.Do("DBSIZE")
	fmt.Printf("size is %d \n", size)
	u := &User{Id: 23, Name: "jianghuan"}
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

func getmsg(req zerver.Request, resp zerver.Response) {
	u := &User{Id: 23, Name: "jianghuan"}
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

func postmsg(req zerver.Request, resp zerver.Response) {
	u := &User{Id: 23, Name: "jianghuan"}
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
	server.Get("/v1/listmsg", listmsg)
	server.Post("/v1/getmsg", getmsg)
	server.Post("/v1/postmsg", postmsg)
	server.Start(nil) // default listen at ":4000"
}
