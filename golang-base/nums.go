package main

import (
	"fmt"
)

//返回byte（char）在数组中的位置
func IndexOfBytes(array []byte, one byte) int {
	for i, entry := range array {
		if entry == one {
			return i
		}
	}
	return -1
}

//返回字符串在数组中的位置
func IndexOfStrings(array []string, one string) int {
	for i, entry := range array {
		if entry == one {
			return i
		}
	}
	return -1
}

//判断字符串是否在数组中
func IsInStrings(array []string, one string) bool {
	return IndexOfStrings(array, one) > -1
}

func main() {
	by := []byte{1, 2, 3, 4, 5}
	str := []string{"abc", "123", "345", "234"}
	fmt.Println(IndexOfBytes(by, 4))
	fmt.Println(IndexOfStrings(str, "345"))
	fmt.Println(IsInStrings(str, "abc"))
}
