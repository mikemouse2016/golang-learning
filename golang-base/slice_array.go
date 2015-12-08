package main
import (
    "fmt"
)
func main() {
    arr1 := [...]int{1, 2, 3}
    //打印初始的指针
    fmt.Printf("the pointer is : %p \n", &arr1)
    printPointer(arr1)

    arr2 := make([]int, 3)
    //打印初始的指针
    fmt.Printf("the pointer is : %p \n", arr2)
    printPointer(arr2)

    arr3 := make(map[int]string)
    //arr := [3]int{1, 2, 3}
    //打印初始的指针
    fmt.Printf("the pointer is : %p \n", arr3)
    printPointer(arr3)

    var array[10]int
    slice := array[0:4]

    fmt.Println(len(slice), cap(slice),len(array), cap(array))

    a := [...]int{1, 2, 3, 4, 5}
    s1 := a[2:4]
    s2 := a[1:5]
    s3 := a[:]
    s4 := a[:4]
    s5 := s2[:]
    fmt.Println(s1,s2,s3,s4,s5)

    s0 := []int{0, 0}
    s1 = append(s0, 2)
    s2 = append(s1, 3, 5, 7)
    s3 = append(s2, s0...)
    fmt.Println(cap(s1), cap(s2), cap(s3))

    b := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
    s := make([]int, 6)
    n1 := copy(s, b[0:])   //n1 == 6, s == []int{0, 1, 2, 3, 4, 5}
    n2 := copy(s, s[2:])
    fmt.Println(n1,n2)

    monthdays := map[string]int{
    "Jan": 31, "Feb": 28, "Mar": 31,
    "Apr": 30, "May": 31, "Jun": 30,
    "Jul": 31, "Aug": 31, "Sep": 30,
    "Oct": 31, "Nov": 30, "Dec": 31,   //逗号是必须的
    }
    //monthdays["Undecim"] = 30   //添加一个月
    fmt.Println(monthdays)
    for months, days := range monthdays {   //键没有使用，因此用_, days
        fmt.Printf("mouth:%s days:%d\n", months, days)
    }
    delete(monthdays, "Jan")
    fmt.Println(monthdays)
}
func printPointer(any interface{}) {
    fmt.Printf("the pointer is : %p \n", &any)
}
