package main

import (
	"fmt"
)

type user struct {
	Name    stringсук
	Surname string
	Phone   string
}

func (u *user) ChangeName(name string) {
	u.Name = name
}

func (u *user) ChangePhone(phone string) {
	u.Phone = phone
}

func Print(user User) {
	fmt.Printf("%+v\n", user)
}

type User interface {
	ChangeName(name string)
	ChangePhone(phone string)
}

func main() {
	user := user{
		Name:    "Bob",
		Surname: "Marley",
		Phone:   "234242343242",
	}
	Print(&user)
	user.ChangeName("Vlad")
	user.ChangePhone("228")
	Print(&user)

}
