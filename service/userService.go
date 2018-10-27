package service

import (
	"AgendaGo/entity"
	"fmt"
)

func hasRegister(username string) bool {
	return false
}

func RegisterUser(username, password, email, phone string) error {

	// check if any information is empty
	if username == "" || password == "" || email == "" || phone == "" {
		return fmt.Errorf("Any of your information should not be empty")
	}

	if hasRegister(username) {
		return fmt.Errorf(username + " has been registered")
	}

	newUser := entity.User{
		Username: username,
		Password: password,
		Email:    email,
		Phone:    phone,
	}

	return nil
}
