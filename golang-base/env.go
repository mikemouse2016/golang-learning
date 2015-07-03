package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("WEB", "http://coolshell.cn")
	fmt.Println(os.Getenv("WEB"))
	fmt.Println(os.Environ())
	for _, env := range os.Environ() {
		e := strings.Split(env, "=")
		fmt.Println(e[0], "=", e[1])
	}
}
