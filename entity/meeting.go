package entity
import (
	"os"
	"fmt"
)
type Meeting struct {
	Title         string   `json:"tile"`
	Sponsor      string   `json:"sponsor"`
	Participators []string `json:"participators"`
	StartDate     string   `json:"startDate"`
	EndDate       string   `json:"endDate"`
}

type meetingModel struct {
	meetings map[string] *Meeting
}

var MeetingModel meetingModel


func (model *meetingModel) AddMeeting(meeting *Meeting) {
	_, ok := model.meetings[meeting.Title]
	if ok == false {
		model.meetings[meeting.Title] = meeting
		fmt.Printf("adding meeting : %s\n",meeting.Title)
	}else{
		fmt.Println(os.Stderr, "Error:%s", "already exists a meeting with the same title")
	}
}

func (model *meetingModel) queryMeeting(title string) (* Meeting,bool){
	_,ok := model.meetings[title]
	if ok == true{
		return model.meetings[title],true
	}else{
		fmt.Println(os.Stderr,"Error:%s","no such meeting")
		return nil,false
	}
}

func (meeting * Meeting) getParticipator() []string{
	return meeting.Participators
}

func (model *meetingModel) deleteMeeting(title string) bool{
	if _,ok := model.meetings[title];ok{
		delete(model.meetings,title)
		return true
	}else
	{
		fmt.Println(os.Stderr,"error:%s\n","no such meeting to delete")
		return false
	}
}
