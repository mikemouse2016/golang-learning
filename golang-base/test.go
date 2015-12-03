
package main

import(
    "fmt"
    "math"
    "math/rand"
    "time"
)

func main(){
 //   fmt.Println(int(3/float32(10)*6))
//    fmt.Println(int(float32(1.23)))
    fmt.Println(math.Trunc(1.734))
    fmt.Println(math.Trunc(-1.234))
    fmt.Println(math.Floor(1.234))
    fmt.Println(math.Floor(-1.234))
    fmt.Println(math.Ceil(1.234))
    fmt.Println(math.Ceil(-1.234))
    fmt.Println(133/2)
    size := 1
    switch {
        case size < 22:
                fmt.Println("22", size)
        case size < 42:
                fmt.Println("42", size)
        default:
                fmt.Println("default", size)
        }
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
fmt.Println(time.Now())

        for i:=0; i<1000000; i++{
             r.Intn(5)
        }
fmt.Println(time.Now())
}

/*
package main
import (
    "fmt"
)
func main() {
    for i := 0; i < 8; i++ {
        for j := 0; j < 8; j++ {
            if (i+j)%2 == 0 {
                fmt.Printf("■")
            } else {
                fmt.Printf("□")
            }
        }
        fmt.Printf("\n")
    }
}
*/
