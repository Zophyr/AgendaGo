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

type meetingDb struct {
	storage
	Data []Meeting `json:"data"`
}

type Meetings struct {
	meetings map[string]*Meeting
}

var AllMeetings Meetings

func (allMeetings *Meetings) AddMeeting(meeting *Meeting) {
	defer allMeetings.dump()
	allMeetings.meetings[meeting.Title] = meeting
}

func (allMeetings *Meetings) DeleteMeeting(meeting *Meeting) {
	defer allMeetings.dump()
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
	if len(curMeetingParticipators) == 0 {
		delete(allMeetings.meetings, meeting.Title)
	}
	allMeetings.meetings[meeting.Title].Participators = curMeetingParticipators
}

// add a participator
func (allMeetings *Meetings) AddParticipatorToMeeting(meeting *Meeting, participator string) {
	curMeetingParticipators := allMeetings.meetings[meeting.Title].Participators
	allMeetings.meetings[meeting.Title].Participators = append(curMeetingParticipators, participator)
}

func (allMeetings *Meetings) QueryMeeting(title string) (*Meeting, bool) {
	defer allMeetings.dump()
	_, err := allMeetings.meetings[title]
	if err != nil {
		return allMeetings.meetings[title], true
	} else {
		fmt.Println(os.Stderr, "Error:%s", "no such meeting")
		return nil, false
	}
}

func (meeting *Meeting) getParticipator() []string {
	return meeting.Participators
}

func (allMeetings *Meetings) load() {
	var meetingDb meetingDb
	allMeetings.storage.load(&meetingDb)
	for index, meeting := range meetingDb.Data {
		allMeetings.meetings[meeting.Title] = &meetingDb.Data[index]
	}
}

func (allMeetings *Meetings) dump() {
	var meetingDb meetingDb
	for _, meeting := range allMeetings.meetings {
		meetingDb.Data = append(meetingDb.Data, *meeting)
	}
	allMeetings.storage.dump(&meetingDb)
}

func (allMeetings *Meetings) Init() { // meeting call this function in the root cmd
	allMeetings.storage.path = "../data/meeting.json"
	allMeetings.meetings = make(map[string]*Meeting)
	allMeetings.load()
}
