package main 
import ( 
    "encoding/json" 
    "fmt" 
    "os" 
)

type demo3 struct {
	Id    int64
	Score int
}
type Group struct {
	Max  int
	Demo []demo3
}

func main ( ) {
    type ColorGroup struct {
        ID     int
        Name   string
        Colors [ ] string
    }
    group := ColorGroup {
        ID :     1 ,
        Name :   "Reds" ,
        Colors : [ ] string { "Crimson" , "Red" , "Ruby" , "Maroon" } ,
    }
    b , err := json. Marshal ( group )
    if err != nil {
        fmt. Println ( "error:" , err )
    }
    os. Stdout . Write ( b )


    g := Group{
			Max:  1,
			Demo: []demo3{{Id : 1, Score: 1}, {Id : 2, Score:2}},
		}
		fmt.Println(g)
		b, err = json.Marshal(g)
		if err != nil {
			fmt.Println("error:", err)
		}
		os.Stdout.Write(b)
}
