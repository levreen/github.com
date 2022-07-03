package main

import "fmt"

// notifier is an interface that defines notification
// type behavior
type notifier interface {
	notify()
}

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements the notifier interface with a pointer receiver.
func (u *user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

func main() {

	// create a value of type user and send a notification.
	u := user{"Bill", "bill@email.com"}

	// values of type user do not implement the interface because pointer
	// receivers don't belong to the method set of a value. 
	sendNotification(&u)
}

func sendNotification(n notifier) {
	n.notify()
}
