package v2

import (
	"fmt"
)

type demo struct {
	a int
	B int
}

type Test struct {
	a int
	B int
}

func Demo() {
	var d demo
	d.a = 10
	d.B = 20
	fmt.Println(d)
}
