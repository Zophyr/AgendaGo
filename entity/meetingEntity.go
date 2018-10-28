package entity

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

//增加与会者
func (model *meetingModel) AddParticipatorToMeeting(meeting *Meeting, participator string) {
	curMeetingParticipators := model.meetings[meeting.Title].Participators
	model.meetings[meeting.Title].Participators = append(curMeetingParticipators, participator)
}
