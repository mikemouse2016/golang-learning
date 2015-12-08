package main
import "fmt"
/*
string是只读的byte slice（和一些额外的属性）。如果你确实需要更新一个字符串，那么使用byte slice，并在需要时把它转换为string类型
*/
func main() {
    x := "text"
//    x[0] = 'T'  error
    fmt.Println(x)

    xbytes := []byte(x)
    xbytes[0] = 'T'
    fmt.Println(string(xbytes)) //prints Text
}
