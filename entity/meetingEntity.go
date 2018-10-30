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

type Meetings struct {
	meetings map[string]*Meeting
}

var AllMeetings Meetings

//找到所有符合条件的会议
//使用filter来实现
func (allMeetings *Meetings) FindBy(cond func(*Meeting) bool) []Meeting {
	result := []Meeting{}
	for _, meeting := range allMeetings.meetings {
		if cond(meeting) {
			result = append(result, *meeting)
		}
	}
	return result
}

func (allMeetings *Meetings) FindByTitle(meetingname string) *Meeting {
	return allMeetings.meetings[meetingname]
}

//删除与会者
func (allMeetings *Meetings) DeleteParticipatorFromMeeting(meeting *Meeting, participator string) {

	curMeetingParticipators := allMeetings.meetings[meeting.Title].Participators
	for i, p := range curMeetingParticipators {
		if p == participator {
			curMeetingParticipators = append(curMeetingParticipators[:i], curMeetingParticipators[i+1:]...)
			break
		}
	}
	allMeetings.meetings[meeting.Title].Participators = curMeetingParticipators
}

//增加与会者
func (allMeetings *Meetings) AddParticipatorToMeeting(meeting *Meeting, participator string) {
	curMeetingParticipators := allMeetings.meetings[meeting.Title].Participators
	allMeetings.meetings[meeting.Title].Participators = append(curMeetingParticipators, participator)
}

func (allMeetings *Meetings) AddMeeting(meeting *Meeting) {
	_, ok := allMeetings.meetings[meeting.Title]
	if ok == false {
		allMeetings.meetings[meeting.Title] = meeting
		fmt.Printf("adding meeting : %s\n", meeting.Title)
	} else {
		fmt.Println(os.Stderr, "Error:%s", "already exists a meeting with the same title")
	}
}

func (allMeetings *Meetings) queryMeeting(title string) (*Meeting, bool) {
	_, ok := allMeetings.meetings[title]
	if ok == true {
		return allMeetings.meetings[title], true
	} else {
		fmt.Println(os.Stderr, "Error:%s", "no such meeting")
		return nil, false
	}
}

func (allMeetings *Meetings) getParticipator() []string {
	return allMeetings.Participators
}

func (allMeetings *Meetings) deleteMeeting(title string) bool {
	if _, ok := allMeetings.meetings[title]; ok {
		delete(allMeetings.meetings, title)
		return true
	} else {
		fmt.Println(os.Stderr, "error:%s\n", "no such meeting to delete")
		return false
	}
}
