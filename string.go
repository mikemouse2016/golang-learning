package main

import(
    "fmt"
)

func main(){

    mystring := "字符串 string"
    byteslice := []byte(mystring)
    //Go 的字符串是UTF-8 编码的，一些字符可能是1、2、3 或者4 个字节结尾
    intslice := []int(mystring)
    //每个int 保存Unicode 编码的指针
    b := []byte{'h','e','l','l','o'}   //复合声明
    s := string(b)
    i := []int{257,1024,65}
    r := string(i)
    /*
    将整数转换到指定的（bit）长度：uint8(int)；
    从浮点数到整数：int(float32)。这会截断浮点数的小数部分；
    其他的类似：float32(int)。
    */
   fmt.Println(mystring, byteslice, intslice) 
}

