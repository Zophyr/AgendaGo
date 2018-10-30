package entity

import (
	"fmt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type Users struct {
	users map[string]*User
}

var AllUsers Users

func (allUsers *Users) AddUser(user *User) {
	allUsers.users[user.Username] = user
}

func (allUsers *Users) DeleteUser(user *User) {
	delete(allUsers.users, user.Username)
}

// use a filter to find appropriate users
func (allUsers *Users) FindBy(cond func(*User) bool) []User {
	result := []User{}
	for _, user := range allUsers.users {
		if cond(user) {
			result = append(result, *user)
		}
	}
	return result
}

func (allUsers *Users) FindByName(username string) []User {
	return allUsers.FindBy(func(user *User) bool {
		return username == user.Username
	})
}

func (allUsers *Users) IsMatchNamePass(username, password string) (bool, error) {
	result := allUsers.FindByName(username)
	if len(result) == 0 {
		return false, fmt.Errorf("The user doesn't exist")
	}
	return result[0].Password == password, nil
}
