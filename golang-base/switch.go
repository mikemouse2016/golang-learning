package main

import(
    "fmt"
)

func main(){
    s1 := "Starting part" 
    s2 := "Ending part"
    s3 := `"hello" 
    'world'`
    fmt.Println(s1+" "+s2+" "+s3)
    var c complex64 = 5+5i;fmt.Printf("Value is: %v\n", c)
    J: for j := 0; j < 5; j++ {
           for i := 0; i < 10; i++ {
              if i > 5 {
                 break J   //现在终止的是j 循环，而不是i 的那个
              }
              println(i)
           }
    }
    switch 0 {
        case 0: // 空的case 体
        case 1:
        fmt.Println("case 1") // 当i == 0 时，f 不会被调用！
        default:
        fmt.Println("default") // 当i 不等于0 或1 时调用
    }
    switch 0 {
        case 0: fallthrough
        case 1:
        fmt.Println("case 1") // 当i == 0 时，f 会被调用！
        default:
        fmt.Println("default") // 当i 不等于0 或1 时调用
    }
    a := [2][2]int{ [2]int{1,2}, [2]int{3,4} }
    fmt.Println(a)
}
