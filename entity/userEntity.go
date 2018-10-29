package entity

import (
	"fmt"
)

type User struct {
	Username string	'json:"username"' 
	Password string	'json:"password"'
	Email    string 'json:"email"'
	Phone    string 'json:"phone"'
}

type UserDB struct{
	Data []User 'json:"Data"'
}

type Users struct {
	storage
	users map[string]*User
}

var AllUsers Users

func (allusers *Users) Init() {      // meeting call this function in the root cmd
	allusers.storage.path = "../data/user.json"
	allusers.meetings = make(map[string]*User)
	allusers.load()
}

func (allUsers *Users) AddUser(user *User) {
	defer allUsers.dump()
	allUsers.users[user.Username] = user
}

func (allUsers *Users) DeleteUser(user *User) {
	defer allUsers.dump()
	delete(allUsers.users, user.Username)
}

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

func (allusers *User) load() {
	var userDB UserDB
	allusers.storage.load(&userDb)
	for index, user := range userDb.Data {
		allusers.users[user.Username] = &userDB.Data[index]
	}
}

func (allusers *User) dump() {
	var userDB UserDB
	for _, user := range allusers.users {
		userDb.Data = append(userDb.Data, *user)
	}
	allusers.storage.dump(&userDB)
}