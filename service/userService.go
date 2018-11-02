package service

import (
	"Agenda/entity"
	"fmt"
)

// user register
func RegisterUser(username, password, email, phone string) error {

	// check if any information is empty
	if username == "" || password == "" || email == "" || phone == "" {
		return fmt.Errorf("One of your information is empty")
	}

	// check if has been register
	if len(entity.AllUsers.FindBy(func(user *entity.User) bool {
		return username == user.Username
	})) != 0 {
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

// user login
func LoginUser(username, password string) error {

	// check if someone has been logged in
	if entity.CurrSession.HasLoggedIn() {
		return fmt.Errorf("You have been logged in")
	}

	// check whether the username and password are correct or not
	isMatch := entity.AllUsers.IsMatchNamePass(username, password)
	if !isMatch {
		return fmt.Errorf("Wrong password")
	}

	// set the current user in the current session
	entity.CurrSession.CurrUser = &entity.AllUsers.FindByName(username)[0]
	return nil
}

// user logout
func LogoutUser() error {

	// check if someone is logged in
	if entity.CurrSession.CurrUser != nil {
		return fmt.Errorf("No one has logged in")
	} else {
		// clear the current user
		entity.CurrSession.CurrUser = nil
		return nil
	}
}

// list all users
func QueryAllUsers() ([]entity.User, error) {

	// check if someone is logged in
	if entity.CurrSession.CurrUser == nil {
		return nil, fmt.Errorf("No one has logged in")
	} else {
		return entity.AllUsers.FindBy(func(user *entity.User) bool {
			return true
		}), nil
	}
}

// delete a user
func DeleteUser() error {

	// check if someone is logged in
	if entity.CurrSession.CurrUser == nil {
		return fmt.Errorf("No one has logged in")
	}

	curUserName := entity.CurrSession.GetCurUserName()

	sponsorMeetings := entity.AllMeetings.FindBy(func(meeting *entity.Meeting) bool {
		// find the meeting whose sponsor is currUser
		if curUserName == meeting.Sponsor {
			return true
		}
		return false
	})

	partiMeetings := entity.AllMeetings.FindBy(func(meeting *entity.Meeting) bool {
		// fint the meeting which currUser participates in
		for _, participator := range meeting.Participators {
			if curUserName == participator {
				return true
			}
		}
		return false
	})

	// delete the sponsor meeting
	for _, meeting := range sponsorMeetings {
		entity.AllMeetings.DeleteMeeting(&meeting)
	}

	// delete the participate meeting
	for _, meeting := range partiMeetings {
		entity.AllMeetings.DeleteParticipator(&meeting, curUserName)
	}

	LogoutUser()
	entity.AllUsers.DeleteUser(&entity.AllUsers.FindByName(curUserName)[0])
	return nil
}
