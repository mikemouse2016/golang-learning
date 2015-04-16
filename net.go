package main
import ( "io/ioutil"; "net/http"; "fmt" )
func main() {
r, err := http.Get("http://www.163.com")
if err != nil { fmt.Println(err); return }
b, err := ioutil.ReadAll(r.Body)
r.Body.Close()
if err == nil { fmt.Printf("%s", string(b)) }
}
