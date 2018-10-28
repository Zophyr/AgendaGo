package entity

import (
	"fmt"
	"os"
)

type Meeting struct {
	Title         string   `json:"tile"`
	Sponsor       string   `json:"sponsor"`
	Participators []string `json:"participators"`
	StartDate     string   `json:"startDate"`
	EndDate       string   `json:"endDate"`
}

type meetingModel struct {
	meetings map[string]*Meeting
}

var MeetingModel meetingModel

func (model *meetingModel) AddMeeting(meeting *Meeting) {
	_, ok := model.meetings[meeting.Title]
	if ok == false {
		model.meetings[meeting.Title] = meeting
		fmt.Printf("adding meeting : %s\n", meeting.Title)
	} else {
		fmt.Println(os.Stderr, "Error:%s", "already exists a meeting with the same title")
	}
}

func (model *meetingModel) queryMeeting(title string) (*Meeting, bool) {
	_, ok := model.meetings[title]
	if ok == true {
		return model.meetings[title], true
	} else {
		fmt.Println(os.Stderr, "Error:%s", "no such meeting")
		return nil, false
	}
}

func (meeting *Meeting) getParticipator() []string {
	return meeting.Participators
}

//找到所有符合条件的会议

//使用filter来实现
func (allMeetings *meetingModel) FindBy(cond func(*Meeting) bool) []Meeting {
	result := []Meeting{}
	for _, meeting := range allMeetings.meetings {
		if cond(meeting) {
			result = append(result, *meeting)
		}
	}
	return result
}

func (model *meetingModel) FindByTitle(meetingname string) *Meeting {
	return model.meetings[meetingname]
}

//删除与会者

func (model *meetingModel) DeleteParticipatorFromMeeting(meeting *Meeting, participator string) {
	//logger.Println("[meetingmodel] try deleting a participator from meeting", meeting.Title)
	curMeetingParticipators := model.meetings[meeting.Title].Participators
	for i, p := range curMeetingParticipators {
		if p == participator {
			curMeetingParticipators = append(curMeetingParticipators[:i], curMeetingParticipators[i+1:]...)
			break
		}
	}
	model.meetings[meeting.Title].Participators = curMeetingParticipators
	//model.dump()
	//logger.Println("[meetingmodel] deleted a participator from meeting", meeting.Title)
}

func (model *meetingModel) deleteMeeting(title string) {
	if _, ok := model.meetings[title]; ok {
		delete(model.meetings, title)
	} else {
		fmt.Println(os.Stderr, "error:%s\n", "no such meeting to delete")
	}
}

//增加与会者
func (model *meetingModel) AddParticipatorToMeeting(meeting *Meeting, participator string) {
	curMeetingParticipators := model.meetings[meeting.Title].Participators
	model.meetings[meeting.Title].Participators = append(curMeetingParticipators, participator)
}
