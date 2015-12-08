package main
import "fmt"
func main() {

    //slice
    slice := []int{1,2,3}
    func(arr []int) {
        arr[0] = 7
        fmt.Println(arr) //prints [7 2 3]
    }(slice)
    fmt.Println(slice) //prints [7 2 3]
    //array
    arr := [3]int{1,2,3}
    func(arr [3]int) {
        arr[0] = 7
        fmt.Println(arr) //prints [7 2 3]
    }(arr)
    fmt.Println(arr) //prints [7 2 3]
}
