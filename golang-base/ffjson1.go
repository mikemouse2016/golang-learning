package main

import (
	"encoding/json"
	"fmt"
        "time"
	"github.com/pquerna/ffjson/ffjson"
)

type Server struct {
	ServerName string `json:"servername"`
	ServerIP   string `json:"serverip"`
}

type Serverslice struct {
	Servers []Server `json:"servers"`
}

func main() {
	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	var buf []byte
        var err error
        now := time.Now().Unix()
	for i := 0; i < 3000000; i++ {
		buf, err = json.Marshal(s)
		if err != nil {
			fmt.Println("json err:", err)
		}
		//fmt.Println(string(b))
	}
        fmt.Println(time.Now().Unix() - now)
	fmt.Println("--------------------------------------------")
	json.Unmarshal(buf, &s)
	fmt.Println(s)
	for _, server := range s.Servers {
		fmt.Println(server)
	}

	fmt.Println("----------------for ffjson--------------------")
	var item = Server{"111", "222"}
        now = time.Now().Unix()
	for i := 0; i < 3000000; i++ {
		buf, err = ffjson.Marshal(s)
		if err != nil {
			fmt.Println("json err:", err)
		}
		//fmt.Println(string(buf))
	}
        fmt.Println(time.Now().Unix() - now)
	ffjson.Unmarshal(buf, &item)
	fmt.Println(item)
}
