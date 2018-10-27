package service

import (
	"AgendaGo/entity"
	"fmt"
)

func RegisterUser(username, password, email, phone string) error {

	// check if any information is empty
	if username == "" || password == "" || email == "" || phone == "" {
		return fmt.Errorf("One of your information is empty")
	}

	if entity.AllUsers.FindBy(func(user *entity.User) bool {
		return username == user.Username
	}) != nil {
		return fmt.Errorf(username + " has been registered")
	}

	newUser := &entity.User{
		Username: username,
		Password: password,
		Email:    email,
		Phone:    phone,
	}

	entity.AllUsers.AddUser(newUser)

	return nil
}
