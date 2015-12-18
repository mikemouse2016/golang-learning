package main
import "fmt"

/*
重新划分一个slice时，新的slice将引用原有slice的数组
*/
func get_err() []byte {
    raw := make([]byte,10000)
    fmt.Println(len(raw),cap(raw),&raw[0]) //prints: 10000 10000 <byte_addr_x>
    return raw[:3]
}

func get_succ() []byte {
    raw := make([]byte,10000)
    fmt.Println(len(raw),cap(raw),&raw[0]) //prints: 10000 10000 <byte_addr_x>
    res := make([]byte,3)
    copy(res,raw[:3])
    return res
}

func main() {
    data1 := get_err()
    fmt.Println(len(data1),cap(data1),&data1[0]) //prints: 3 10000 <byte_addr_x>
    fmt.Println("==============")
    data2 := get_succ()
    fmt.Println(len(data2),cap(data2),&data2[0]) //prints: 3 10000 <byte_addr_x>
}
