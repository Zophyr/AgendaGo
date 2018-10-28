package service

import (
	"entity"
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


//err := service.DeleteFromMeeting(title) 
func deleteFromMeetingByTitle(title string) error{
	if meeting,err := model.queryMeeting(title);err{
		model.deleteMeeting(title)
		return nil
	}else
	{
		return error("no meeting to be deleted")
	}
}

//service.queryMeeting(title).getParticipator().empty()
func queryMeeting(title string) *entity.meetings{
	for k,v := model.meetings{
		if k == string{
			return v
		}
	}
	return nil
}

func existsInParticipator(participators []string,userName string) bool{
	for s:= range participators{
		if userName == s{
			return true
		}
	}
	return false
}

// need global variable userName indicates the current login user
func quitMeeting(title string) error{
	for k,v := model.meetings{
		if k == title && existsInParticipator(v.participators,userName){
			for i:=0;i<len(v.participators);i++{
				if(v.participators[i]==userName){
					v.participators = append(v.participators[:i]+v.participators[i+1:])
					return nil
				}
			}
		}
	}
	return error("doesnt find it")
}

func DeleteAllMeetings(title string) error{
	for k,v : model,meetings{
		delete(model.meetings,k)
	}
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

func LogoutUser() error {

	if entity.CurrSession.CurrUser != nil {
		return fmt.Errorf("No one has logged in")
	} else {
		entity.CurrSession.CurrUser = nil
		return nil
	}
}

func QueryAllUsers() ([]entity.User, error) {
	if entity.CurrSession.CurrUser != nil {
		return nil, fmt.Errorf("No one has logged in")
	} else {
		return entity.AllUsers.FindBy(func(user *entity.User) bool {
			return true
		}), nil
	}
}
