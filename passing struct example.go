package main

import "fmt"

// user defines a user in the program.
type user struct {
	name   string
	email  string
	status string
}

func debug(u user) {
	fmt.Printf("%+v\n", u)
}

// notify implements a method with a value receiver.
func (u user) notify() {
	fmt.Printf("User: Sending User Email To %s<%s>\n", u.name, u.email)
}

// changeEmail implements a method with a pointer receiver.
func (u *user) changeEmail(email string) {
	u.email = email
}

func (u *user) changeName(name string) {
	u.name = name
}

func (u *user) Send() {
	u.notify()
}

// main is the entry point for the application.
func main() {
	// Pointers of type user can also be used to methods
	// declared with a value receiver.
	john := &user{"John", "john@exemple.com", "enabled"}

	john.changeName("John Smith")
	john.changeEmail("john@gmail.com")
	john.notify()

	Receive(john)
	debug(user{"Willy", "Willy@exemple.com", "enabled"})
}

func Receive(user interface {
	changeName(s string)
	changeEmail(s string)
	Send()
}) {
	user.changeName("Bill")
	user.changeEmail("bill@billy-exemple.com")
	user.Send()
}
