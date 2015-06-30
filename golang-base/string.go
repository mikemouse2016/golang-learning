package main

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

//将一个字节数组转换成utf8字符串
func Utf8(bs []byte) (str string, err error) {
	if utf8.FullRune(bs) {
		str = string(bs)
		return
	}
	//错误
	err = errors.New("fail to decode to utf8")
	str = ""
	return
}

//判断一个字符串是不是空字符串
func IsSpace(c byte) bool {
	if c >= 0x00 && c <= 0x20 {
		return true
	}
	return false
}

//判断一个字符串是不是空白串，即（0x00 - 0x20 之内的字符均为空白字符）
func IsBlank(str string) bool {
	for i := 0; i < len(str); i++ {
		b := str[i]
		if !IsSpace(b) {
			return false
		}
	}
	return true
}

//如果输入的字符串为空串，那么返回默认字符串
func SBlank(str, dft_str string) string {
	if IsBlank(str) {
		return dft_str
	}
	return str
}

//将字符串转换成整数，如果转换失败，采用默认值
func ToInt(str string, dft int) int {
	re, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return dft
	}
	return re
}

//拆分字符串数组，如果数组元素为空白，忽略，否则Trim空白
func SplitIgnoreBlank(str, sep string) []string {
	ss := strings.Split(str, sep)
	size := len(ss)
	re := make([]string, size)
	for i := 0; i < size; i++ {
		str := Trim(ss[i])
		if len(str) > 0 {
			re = append(re, str)
		}
	}
	return re

}

//去掉一个字符串左右的空白串，即（0x00 - 0x20 之内的字符均为空白字符）
//与strings.TrimSpace功能一致
func Trim(str string) string {
	size := len(str)
	if size <= 0 {
		return s
	}
	
	for l := 0 ; l < size l++ {
		b := s[l]
		if !IsSpace(b) {
			break
		}
	}
	
	for r := size-1; r >= 1; r-- {
		b:=s[r]
		if !IsSpace(b){
			break
		}
	}
	return string(s[l:r+1])
}

func TrimBytes(bs []byte) string {
	r := len(bs)-1 
	if r <= 0 {
		return string(bs)
	}
	for l:=0; l<r-1; l++ {
		b:=bs[l]
		if !IsSpace(b){
			break
		}
	}
	for ; r >= l; r-- {
		b := bs[r]
		if !IsSpace(b) {
			break
		}
	}
	return string(bs[l, r+1])
}


func main() {
	mystring := "字符串 string"
	byteslice := []byte(mystring)

	//Go 的字符串是UTF-8 编码的，一些字符可能是1、2、3 或者4 个字节结尾
	intslice := []int(mystring)
	//每个int 保存Unicode 编码的指针
	b := []byte{'h', 'e', 'l', 'l', 'o'} //复合声明
	s := string(b)
	i := []int{257, 1024, 65}
	r := string(i)
	/*
	   将整数转换到指定的（bit）长度：uint8(int)；
	   从浮点数到整数：int(float32)。这会截断浮点数的小数部分；
	   其他的类似：float32(int)。
	*/
	fmt.Println(mystring, byteslice, intslice)
}
