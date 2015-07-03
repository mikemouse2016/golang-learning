package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("ping", "127.0.0.1", "-c", "2")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Command Error!", err.Error())
		return
	}
	fmt.Println(string(out))

	cmd = exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input")
	var out1 bytes.Buffer
	cmd.Stdout = &out1
	err = cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("in all caps: %q\n", out1.String())
}
