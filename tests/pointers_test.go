package main

import "testing"

func TestUser_ChangeName(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		user := UserFactory()
		user.ChangeName("Loh")
		if user.Name != "Loh" {
			t.Error("Username should be Loh")
		}
	})
}

func TestUser_ChangePhone(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		user := UserFactory()
		user.ChangePhone("223")
		if user.Phone != "223" {
			t.Error("Phone should be 223")
		}
	})
}

func UserFactory() user {
	return user{
		Name:    "Bob",
		Surname: "Marley",
		Phone:   "234242343242",
	}
}
