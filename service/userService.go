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

func LogoutUser() error {

	if entity.CurrSession.CurrUser != nil {
		return fmt.Errorf("No one has logged in")
	} else {
		entity.CurrSession.CurrUser = nil
		return nil
	}
}

func QueryAllUsers() ([]entity.User, error) {
	if entity.CurrSession.CurrUser == nil {
		return nil, fmt.Errorf("No one has logged in")
	} else {
		return entity.AllUsers.FindBy(func(user *entity.User) bool {
			return true
		}), nil
	}
}

//删除用户

func DeleteUser() error {
	//检验是否登陆
	if entity.CurrSession.CurrUser == nil {
		return fmt.Errorf("No one has logged in")
	}

	curUserName := entity.CurrSession.GetCurUser()

	meetings := entity.MeetingModel.FindBy(func(m *meeting) bool {
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
