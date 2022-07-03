package main

import (
	"fmt"
)

type user struct {
	name    string
	surname string
}

func main() {

	// declare and make a map that stores values
	// of type user with a key of type string
	users := make(map[string]user)

	// add key/value pairs to the map.
	users["Roy"] = user{"Rob", "Roy"}
	users["Ford"] = user{"Henry", "Ford"}
	users["Mouse"] = user{"Mickey", "Mouse"}
	users["Jackson"] = user{"Michael", "Jackson"}

	// read the value of a specific key
	mouse := users["Mouse"]

	fmt.Printf("%#v\n", mouse)

	// replace the value at the Mouse key
	users["Mouse"] = user{"Jerry", "Mouse"} 

	// read the mouse key again.
	fmt.Println("%#v\n", users["Mouse"])

	// delete the value at a specific key
	delete(users, "Roy")

	// check the length of the map. There are only 3 elements.
	fmt.Println(len(users))

	for i := range users {
		fmt.Println(users[i])
	}

}
