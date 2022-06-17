package main

import (
	"fmt"
)

type user struct {
	name    string
	email   string
	isAdmin bool
}

func (u user) isAdminUser() bool {

	return u.isAdmin
}

func (u user) GetName() string {

	return u.name
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type NameLs interface {
	GetName() string
	notify()
}

func main() {
	u := user{name: "salem", email: "salem,.com"}
	sendNotification(&u)
}


func GetEntName(ent NameLs) string {

	return ent.GetName()
}


func sendNotification(ent NameLs) {

	ent.notify()
}