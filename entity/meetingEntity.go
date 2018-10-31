package entity

type Meeting struct {
	Title         string   `json:"tile"`
	Sponsor       string   `json:"sponsor"`
	Participators []string `json:"participators"`
	StartTime     string   `json:"startTime"`
	EndTime       string   `json:"endTime"`
}

type MeetingDB struct {
	Data []Meeting `json:"data"`
}

type Meetings struct {
	storage
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
func (allMeetings *Meetings) DeleteParticipator(meeting *Meeting, participator string) {

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

// add a participator to a meeting
func (allMeetings *Meetings) AddParticipatorToMeeting(meeting *Meeting, participator string) {
	curMeetingParticipators := allMeetings.meetings[meeting.Title].Participators
	allMeetings.meetings[meeting.Title].Participators = append(curMeetingParticipators, participator)
}

func (allMeetings *Meetings) load() {
	var meetingDB MeetingDB
	allMeetings.storage.load(&meetingDB)
	for index, meeting := range meetingDB.Data {
		allMeetings.meetings[meeting.Title] = &meetingDB.Data[index]
	}
}

func (allMeetings *Meetings) dump() {
	var meetingDb MeetingDB
	for _, meeting := range allMeetings.meetings {
		meetingDb.Data = append(meetingDb.Data, *meeting)
	}
	allMeetings.storage.dump(&meetingDb)
}

func (allMeetings *Meetings) Init(path string) { // meeting call this function in the root cmd
	allMeetings.storage.path = "../data/meeting.json"
	allMeetings.meetings = make(map[string]*Meeting)
	allMeetings.load()
}

func init() {
	addModel(&AllMeetings, "meeting_data")
}
