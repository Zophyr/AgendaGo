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

func LoginUser(username, password string) error {

	if entity.CurrSession.HasLoggedIn() {
		return fmt.Errorf("You have been logged in")
	}

	isMatch, err := entity.AllUsers.IsMatchNamePass(username, password)
	if err != nil {
		return err
	}
	if !isMatch {
		return fmt.Errorf("Wrong password")
	}

	entity.CurrSession.CurrUser = &entity.AllUsers.FindByName(username)[0]
	return nil
}
