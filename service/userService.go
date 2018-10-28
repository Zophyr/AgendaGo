package service

import (
	"entity"
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