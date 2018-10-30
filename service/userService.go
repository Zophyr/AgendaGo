package service

import (
	"AgendaGo/entity"
	"fmt"
)

// user register
func RegisterUser(username, password, email, phone string) error {

	// check if any information is empty
	if username == "" || password == "" || email == "" || phone == "" {
		return fmt.Errorf("One of your information is empty")
	}

	// check if has been register
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

// user login
func LoginUser(username, password string) error {

	// check if someone has been logged in
	if entity.CurrSession.HasLoggedIn() {
		return fmt.Errorf("You have been logged in")
	}

	// check whether the username and password are correct or not
	isMatch, err := entity.AllUsers.IsMatchNamePass(username, password)
	if err != nil {
		return err
	}
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

	meetings := entity.AllMeetings.FindBy(func(m *meeting) bool {
		//查找删除用户作为sponsor的会议
		if curUserName == m.Sponsor {
			return true
		}
		//查找作为participator
		for _, participator := range m.Participators {
			if curUserName == participator {
				return true
			}
		}
		return false
	})

	//在会议中删除该用户
	for _, meeting := range meetings {
		err := DeleteFromMeeting(meeting.Title)
		if err != nil {
			return err
		}
	}

	//成功全部处理
	LogoutUser()
	entity.AllUsers.DeleteUser(&entity.AllUsers.FindByName(curUserName)[0])
	return nil
}


// list 
func QueryMeeting(title string) *entity.Meetings{
	for k,v := entity.AllMeetings.meetings{
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
	for k,v := model,meetings{
		delete(model.meetings,k)
	}
	return nil
}
