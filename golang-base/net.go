package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
//	"time"
        "strings"
)

func main() {
//	for i := 0; i < 1000; i++ {
//		time.Sleep(time.Second * 2)
		r, err := http.Get("http://localhost/export.txt")
                if r != nil {
		    defer r.Body.Close()
                }
		if err != nil {
		    fmt.Println(err)
		    return
		}
//                fmt.Print(i)
//		b, err := ioutil.ReadAll(r.Body)
		b, err := ioutil.ReadAll(r.Body)
//		if err == nil {
//			fmt.Printf("%s", string(b))
//		}
                fields := strings.Fields(string(b))
                for key,value := range fields{

                    fmt.Println(key, value)
                }
//	}

}
