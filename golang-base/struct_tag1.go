package main

import (
     "fmt"
     "reflect"
 )

 func main() {
     type S struct {
         F string `species:"gopher" color:"blue"`
     }

     s := S{}
     st := reflect.TypeOf(s)
     field := st.Field(0)
     fmt.Println("species:", field.Tag.Get("species"), "color:", field.Tag.Get("color"))
}
