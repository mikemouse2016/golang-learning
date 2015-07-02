package main

import (
	"fmt"
	"time"
)

func main() {
	//Ticker持续通知
	ticker := time.NewTicker(time.Second)

	for t := range ticker.C {
		fmt.Println("Tick at", t)
	}
}
