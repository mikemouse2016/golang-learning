package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Second * 2)
		r, err := http.Get("http://www.thel-service.com/")
		if err != nil {
			fmt.Println(err)
			return
		}
                fmt.Print(i)
		b, err := ioutil.ReadAll(r.Body)
		r.Body.Close()
		if err == nil {
			fmt.Printf("%s", string(b))
		}
	}

}
