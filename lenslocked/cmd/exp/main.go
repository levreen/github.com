package main

import (
	"fmt"
	"html/template"
	"os"
)

// User for the template
type User struct {
	Name string
	Age  int
	Meta UserMeta
}

// UserMeta for nested struct
type UserMeta struct {
	Visits int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	// Anonymous struct
	// user := struct {
	// 	Name string
	// }{
	// 	Name: "Levi Ocampo",
	// }

	user := User{
		Name: "Levi Ocampo",
		Age:  100,
		Meta: UserMeta{
			Visits: 4,
		},
	}

	fmt.Println(user.Meta.Visits)

	err1 := t.Execute(os.Stdout, user)
	if err1 != nil {
		panic(err1)
	}
}
