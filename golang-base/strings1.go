package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	field1 := "*"
	field2 := "1-19/32"
	field3 := "  foo1;bar2,baz3..."
	ranges := strings.FieldsFunc(field1, func(r rune) bool { return r == ',' })
	fmt.Println(len(ranges))
	for _, expr := range ranges {
		fmt.Println(expr)
	}
	ranges = strings.FieldsFunc(field2, func(r rune) bool { return r == ',' })
	fmt.Println(len(ranges))
	for _, expr := range ranges {
		fmt.Println(expr)
	}

	//
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q", strings.FieldsFunc(field3, f))
}
