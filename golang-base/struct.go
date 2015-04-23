package main

import (
	"fmt"
)

type User struct {
	Name  string
	Email string
}

type Admin struct {
	User
	Name string
}

func (this *User) Hello() {
	fmt.Println(this.Name)
}

/*
func (this *Admin) Hello() {
	fmt.Println(this.Name)
}
*/
func (this *Admin) Notify() {
	//Name := "this.hello"
	//this.Name = Name
	this.Hello()
}

func main() {
	a := &Admin{Name: "name"} //User: {Email: "user.email", Name: "user.Name"}
	a.User.Email = "user.email"
	a.User.Name = "user.name"
	a.Notify()
}
