package service

import (
	"AgendaGo/entity"
	"fmt"
	"time"
)

type meeting = entity.Meeting

const timeFormat = "2018-10-28/13:08:22"

func validateFreeTime(startTime string, endTime string, meetings []entity.Meeting) error {
	for _, meeting := range meetings {
		if endTime > meeting.StartTime && meeting.EndTime > startTime {
			return fmt.Errorf("Conflit with meeting %s", meeting.Title)
		}
	}
	return nil
}

func validateNewMeeting(meeting *entity.Meeting) error {

	if len(meeting.StartTime) == 0 {
		return fmt.Errorf("Start time is empty")
	}
	_, err := time.Parse(timeFormat, meeting.StartTime)

	if err != nil {
		return fmt.Errorf("StartTime format is illegal")
	}

	if len(meeting.EndTime) == 0 {
		return fmt.Errorf("End time is empty")
	}
	_, err = time.Parse(timeFormat, meeting.EndTime)

	if err != nil {
		return fmt.Errorf("EndTime format is illegal")
	}

	if meeting.StartTime > meeting.EndTime {
		return fmt.Errorf("StartTime is after EndTime")
	}

	meetings := entity.AllMeetings.FindBy(func(m *entity.Meeting) bool {
		if m.Sponsor == meeting.Sponsor {
			return true
		}
		// todo check
		for _, participator := range m.Participators {
			if participator == meeting.Sponsor {
				return true
			}
		}
		return false
	})

	for _, participatorName := range meeting.Participators {
		meetings = append(meetings, entity.AllMeetings.FindBy(func(m *entity.Meeting) bool {
			if m.Sponsor == participatorName {
				return true
			}
			// todo check
			for _, participator := range m.Participators {
				if participator == participatorName {
					return true
				}
			}
			return false
		})...)
	}

	return validateFreeTime(meeting.StartTime, meeting.EndTime, meetings)

}

func AddMeetingToCurrSession(title string, participatorName []string, startTime string, endTime string) error {

	// check if someone has logged in
	if !entity.CurrSession.HasLoggedIn() {
		return fmt.Errorf("You have not logged in")
	}

	if len(title) == 0 {
		return fmt.Errorf("Title should not be empty")
	}

	if len(participatorName) == 0 {
		return fmt.Errorf("Meeting must have participator")
	}

	// check if there exists meetings with same name
	if len(entity.AllMeetings.FindByTitle(title)) != 0 {
		return fmt.Errorf("The name same as %s has been added", title)
	}

	sponsorName := entity.CurrSession.GetCurUserName()

	// check if the participators exist
	for _, partiName := range participatorName {
		if len(entity.AllUsers.FindByName(partiName)) == 0 {
			return fmt.Errorf("User %s doesn't exist", partiName)
		}
		if partiName == sponsorName {
			return fmt.Errorf("The sponsor can't be the participator")
		}
	}

	newMeeting := &meeting{
		Title:         title,
		Sponsor:       sponsorName,
		Participators: participatorName,
		StartTime:     startTime,
		EndTime:       endTime,
	}

	if err := validateNewMeeting(newMeeting); err != nil {
		return err
	}

	entity.AllMeetings.AddMeeting(newMeeting)
	return nil
}

func QuitFromMeeting(title string) error {

	// check if someone has logged in
	if !entity.CurrSession.HasLoggedIn() {
		return fmt.Errorf("You have not logged in")
	}

	meeting := entity.AllMeetings.FindByTitle(title)
	if len(meeting) == 0 {
		return fmt.Errorf("Meeting %s doesn't exist", title)
	}

	if meeting[0].Sponsor != entity.CurrSession.GetCurUserName() {
		entity.AllMeetings.DeleteParticipator(&meeting[0], entity.CurrSession.GetCurUserName())
		return nil
	} else {
		entity.AllMeetings.DeleteMeeting(&meeting[0])
	}
	return nil
}

// Delete the meeting whose name is title
// The logged user should be the sponsor
func DeleteMeetingByTitle(title string) error {

	// check if someone has logged in
	if !entity.CurrSession.HasLoggedIn() {
		return fmt.Errorf("You have not logged in")
	}

	meeting := entity.AllMeetings.FindByTitle(title)
	if len(meeting) == 0 {
		return fmt.Errorf("Meeting %s doesn't exist", title)
	}
	if meeting[0].Sponsor != entity.CurrSession.GetCurUserName() {
		return fmt.Errorf("You are not the sponsor of meeting %s", title)
	}
	entity.AllMeetings.DeleteMeeting(&meeting[0])
	return nil
}

func DeleteParticipatorFromMeeting(title string, participatorNames []string) error {

	if !entity.CurrSession.HasLoggedIn() {
		return fmt.Errorf("No one has logged in")
	}

	curMeeting := entity.AllMeetings.FindByTitle(title)
	if len(curMeeting) == 0 {
		return fmt.Errorf("Meeting %s doesn't exist", title)
	}

	for _, participatorName := range participatorNames {
		participator := entity.AllUsers.FindByName(participatorName)
		if len(participator) == 0 {
			return fmt.Errorf("User %s doesn't exist", participatorName)
		}

		flag := false
		for _, curParticipator := range curMeeting[0].Participators {
			if curParticipator == participatorName {
				flag = true
			}
		}

		if !flag {
			return fmt.Errorf("Participator %s is not in the meeting %s", participatorName, title)
		}

		entity.AllMeetings.DeleteParticipator(&curMeeting[0], participatorName)
	}
	return nil
}

func AddParticipatorToMeeting(title string, participatorNames []string) error {

	if !entity.CurrSession.HasLoggedIn() {
		return fmt.Errorf("No one has logged in")
	}

	targetMeeting := entity.AllMeetings.FindByTitle(title)
	if len(targetMeeting) == 0 {
		return fmt.Errorf("Meeting %s doesn't exist", title)
	}

	for _, participatorName := range participatorNames {
		participator := entity.AllUsers.FindByName(participatorName)
		if len(participator) == 0 {
			return fmt.Errorf("User %s doesn't exist", participatorName)
		}

		meetings := entity.AllMeetings.FindBy(func(meeting *entity.Meeting) bool {
			if meeting.Sponsor == participatorName {
				return true
			}
			for _, pName := range meeting.Participators {
				if pName == participatorName {
					return true
				}
			}
			return false
		})

		return validateFreeTime(targetMeeting[0].StartTime, targetMeeting[0].EndTime, meetings)
	}

	for _, participatorName := range participatorNames {
		entity.AllMeetings.AddParticipatorToMeeting(&targetMeeting[0], participatorName)
	}

	return nil
}

func QueryMeeting(startTime, endTime string) (*entity.Meeting, error) {
	return nil, nil
}

func ClearAllMeeting() error {

	if !entity.CurrSession.HasLoggedIn() {
		return fmt.Errorf("No one has logged in")
	}

	meetings := entity.AllMeetings.FindBy(func(meeting *entity.Meeting) bool {
		return entity.CurrSession.GetCurUserName() == meeting.Sponsor
	})

	for _, meeting := range meetings {
		entity.AllMeetings.DeleteMeeting(&meeting)
	}

	return nil
}
