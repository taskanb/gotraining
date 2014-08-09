// NEED PLAYGROUND

// Sample program to show how to use an interface in Go. In this case,
// a pointer is used to support the interface call.
package main

import (
	"fmt"
)

type (
	// notifier is an interface that defined notification
	// type behavior.
	notifier interface {
		notify()
	}

	// user defines a user in the program.
	user struct {
		name  string
		email string
	}
)

// notify implements a method that can be called via
// a value of type user.
func (u *user) notify() {
	fmt.Printf("user: Sending user Email To %s<%s>\n",
		u.name,
		u.email)
}

// main is the entry point for the application.
func main() {
	// Create two values of type user.
	user1 := user{"Bill", "bill@email.com"}
	user2 := &user{"Jill", "jill@email.com"}

	// Pass a pointer of the values to support the interface.
	sendNotification(&user1)
	sendNotification(user2)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(notify notifier) {
	notify.notify()
}
