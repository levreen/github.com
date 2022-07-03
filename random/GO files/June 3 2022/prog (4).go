package main

import (
	"fmt"
)

type user struct {
	likes int
}

func main() {

	// declare a slice of 3 users
	users := make([]user, 3)

	// share the user at index 1
	shareUser := &users[1]

	// add a like for the user that was shared
	shareUser.likes++

	// display the number of likes for all users.

	for i := range users {
		fmt.Printf("User: %d Likes: %d\n", i, users[i].likes)
	}

	// appending a new user introduces an issue
	users = append(users, user{}) // ano to user{}
	users = append(users, user{}) // mag-add nga syang isa pang user

	// add more likes for the user that was shared
	users[3].likes++
	// display the number of likes for all users.

	for i := range users {
		fmt.Printf("User: %d Likes: %d\n", i, users[i].likes)
	}

}
