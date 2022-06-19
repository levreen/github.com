package main

import (
	"html/template"
	"os"
)

// User for the template
type User struct {
	Name string
	Bio  string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	user := User{
		Name: "Levi Ocampo",
		Bio:  "<script>alert('Haha, youve been h4x0r3d!');</script>",
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
