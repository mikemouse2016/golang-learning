package main

import (
	"flag"
	"fmt"
)

func main() {

	//第一个参数是“参数名”，第二个是“默认值”，第三个是“说明”。返回的是指针
	host := flag.String("host", "coolshell", "a host name")

	port := flag.Int("port", 80, "a port number")

	debug := flag.Bool("d", false, "enable/disable debug mode")

	//正式开始Parse命令行参数
	flag.Parse()

	fmt.Println("host:", *host)

	fmt.Println("port:", *port)

	fmt.Println("debug:", *debug)
}

/*
#如果没有指定参数名，则使用默认值
$ go run flagtest.go
host: coolshell.cn
port: 80
debug: false

#指定了参数名后的情况
$ go run flagtest.go -host=localhost -port=22 -d
host: localhost
port: 22
debug: true

#用法出错了（如：使用了不支持的参数，参数没有=）
$ go build flagtest.go
$ ./flagtest -debug -host localhost -port=22
flag provided but not defined: -debug
Usage of flagtest:
  -d=false: enable/disable debug mode
  -host="coolshell.cn": a host name
  -port=80: a port number
exit status 2

*/
