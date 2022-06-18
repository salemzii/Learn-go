package demo

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


go func (){
	defer wg.Done()

	for count := 0; count < 3; count++ {
		for char := 'a'; char < 'a' +26; char++ {
			fmt.Printf("%c ", char)
		}
	} 
}()