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

func (allMeetings *Meetings) AddMeeting(meeting *Meeting) {
	allMeetings.meetings[meeting.Title] = meeting
}

func (allMeetings *Meetings) DeleteMeeting(meeting *Meeting) {
	delete(allMeetings.meetings, meeting.Title)
}

// use a filter to find appropriate meetings
func (allMeetings *Meetings) FindBy(cond func(*Meeting) bool) []Meeting {
	result := []Meeting{}
	for _, meeting := range allMeetings.meetings {
		if cond(meeting) {
			result = append(result, *meeting)
		}
	}
	return result
}

func (allMeetings *Meetings) FindByTitle(title string) []Meeting {
	return allMeetings.FindBy(func(meeting *Meeting) bool {
		return title == meeting.Title
	})
}

// delete a participator
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

// add a participator
func (allMeetings *Meetings) AddParticipatorToMeeting(meeting *Meeting, participator string) {
	curMeetingParticipators := allMeetings.meetings[meeting.Title].Participators
	allMeetings.meetings[meeting.Title].Participators = append(curMeetingParticipators, participator)
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

func (meeting *Meeting) getParticipator() []string {
	return meeting.Participators
}
