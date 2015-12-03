package main

import(
	"fmt"
	"github.com/coocood/jas"
	"net/http"
)
type Age struct{
	a int
}
type Hello struct {
	a Age
	Id int
	Name string
}

func (*Hello) Get (ctx *jas.Context) { // `GET /v1/hello`
    //ctx.Data = "hello world"
	ctx.ParseForm()
	fmt.Println(ctx.Data)
	fmt.Println(ctx.Host)
	fmt.Println(ctx.RequestURI)
	fmt.Println(ctx.Method)
	fmt.Println(ctx.RemoteAddr)
	fmt.Println(ctx.Form)
	fmt.Println(ctx.Form.Get("id"))
	fmt.Println(ctx.Form.Get("a"))
    a := Hello{Id:1,Name:"jianghuan"}
	a.a.a = 33
	ctx.Data = a;
       fmt.Println(ctx.Id)
    //response: `{"data":"hello world","error":null}`
}

func (*Hello) Post (ctx *jas.Context) { // `GET /v1/hello`
    //ctx.Data = "hello world"
	ctx.ParseForm()
	fmt.Println(ctx.Data)
	fmt.Println(ctx.Host)
	fmt.Println(ctx.RequestURI)
	fmt.Println(ctx.Method)
	fmt.Println(ctx.RemoteAddr)
	fmt.Println(ctx.Form)
	fmt.Println(ctx.Form.Get("id"))
	fmt.Println(ctx.Form.Get("a"))
	ctx.Data = Hello{Id:1,Name:"jianghuan"}
    //response: `{"data":"hello world","error":null}`
}

/*
func (*Users) Photo (ctx *jas.Context) {
    // will stop execution and response `{"data":null,"error":"nameInvalid"} if "name" parameter is not given..
    name := ctx.RequireString("name")
    age := ctx.RequirePositiveInt("age")
    grade, err := ctx.FindPositiveInt("grade")

    // 6, 60 is the min and max length, error message can be "passwordTooShort" or "passwordTooLong"
    password := ctx.RequireStringLen(6, 60, "password")

    // emailRegexp is a *regexp.Regexp instance.error message would be "emailInvalid"
    email := ctx.RequireStringMatch(emailRegexp, "email")
    _, _, _, _, _, _ = name, age, grade, err,password, email
}
*/

type Match struct {}

//func (*Users) Gap() string {
//		return ":name"
//}

func (*Match) Photo (ctx *jas.Context) {// `GET /users/:name/photo`
	// suppose the request path is `/users/john/photo`
        name := ctx.GapSegment("") // "john"
        _ = name
}

func (*Match) UserId (ctx *jas.Context) { //`GET /users/:name/photo/:id`
	// suppose the request path is `/users/john/photo/7`
	id := ctx.Id // 7
        fmt.Println(321)
        fmt.Println(id)
        _ = id
}

func main () {
    router := jas.NewRouter(new(Hello), new(Match))
    router.BasePath = "/v1/"
    fmt.Println(router.HandledPaths(true))
    //output: `GET /v1/hello`
    http.Handle(router.BasePath, router)
    http.ListenAndServe(":8080", nil)
}
