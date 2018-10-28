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
