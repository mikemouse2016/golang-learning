package main

import(
    "fmt"
    "math"
    "time"
)

var z float64 = 1

func algo(like int, unlike int)float64{
    n := like+unlike
    if n==0 {
        return 0
    }
    phat := float64(like) / float64(n)
    return (phat+z*z/(2*float64(n))-z*math.Sqrt((phat*(1-phat)+z*z/(4*float64(n)))/float64(n)))/(1+z*z/float64(n))

}

func main(){
/*
    fmt.Println(time.Now())
    for i := 1; i< 1000000; i++{
        fmt.Println(algo(i+100,i+102))
    }
    fmt.Println(time.Now())
*/
//    fmt.Println(algo(89, 213))
    fmt.Println(time.Now())

    for i:= 0.00001; i<1; i+=0.00001{
        math.Sqrt(math.Sqrt(i))
    }
    fmt.Println(time.Now())

    fmt.Println(math.Sqrt(0.02))
    fmt.Println(math.Sqrt(0.021))
    fmt.Println(math.Sqrt(0.03))
    fmt.Println(math.Sqrt(0.2))
    fmt.Println(math.Sqrt(0.4))

    fmt.Println(math.Cbrt(0.02))
    fmt.Println(math.Cbrt(0.021))
    fmt.Println(math.Cbrt(0.03))
    fmt.Println(math.Cbrt(0.2))
    fmt.Println(math.Cbrt(0.4))
}
