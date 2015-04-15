package main

import(
"fmt"
)
func main(){
b := []byte{'h','e','l','l','o'}   //复合声明
s := string(b)
i := []int{257,1024,65}
//r := string(i)
fmt.Printf("%s\n",s, i)
}
