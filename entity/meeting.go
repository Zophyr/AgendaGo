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

type meetingDb struct {
	Data []Meeting `json:"data"`
}

type meetingModel struct {
	storage
	meetings map[string] *Meeting
}

var MeetingModel meetingModel

func (model *meetingModel) Init() {      // meeting call this function in the root cmd
	model.storage.path = "../data/meeting.json"
	model.meetings = make(map[string]*Meeting)
	model.load()
}

func (model *meetingModel) AddMeeting(meeting *Meeting) {
	defer model.dump()
	_, ok := model.meetings[meeting.Title]
	if ok == false {
		model.meetings[meeting.Title] = meeting
		fmt.Printf("adding meeting : %s\n",meeting.Title)
	}else{
		fmt.Println(os.Stderr, "Error:%s", "already exists a meeting with the same title")
	}
}

func (model *meetingModel) queryMeeting(title string) (* Meeting,bool){
	defer model.dump()
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
	defer model.dump()
	if _,ok := model.meetings[title];ok{
		delete(model.meetings,title)
		return true
	}else
	{
		fmt.Println(os.Stderr,"error:%s\n","no such meeting to delete")
		return false
	}
}


func (model *meetingModel) load() {
	var meetingDb meetingDb
	model.storage.load(&meetingDb)
	for index, meeting := range meetingDb.Data {
		model.meetings[meeting.Title] = &meetingDb.Data[index]
	}
}

func (model *meetingModel) dump() {
	var meetingDb meetingDb
	for _, meeting := range model.meetings {
		meetingDb.Data = append(meetingDb.Data, *meeting)
	}
	model.storage.dump(&meetingDb)
}