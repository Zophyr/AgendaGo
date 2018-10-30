package service

import (
	"AgendaGo/entity"
	"fmt"
	"time"
)

type meeting = entity.Meeting

const timeFormat = "2018-10-28/13:08:22"

func validateNewMeeting(meeting *meeting) error {

	if len(meeting.StartDate) == 0 {
		return fmt.Errorf("Start time is empty")
	}
	_, err := time.Parse(timeFormat, meeting.StartDate)

	if err != nil {
		return fmt.Errorf("StartTime format is illegal")
	}

	if len(meeting.EndDate) == 0 {
		return fmt.Errorf("End time is empty")
	}
	_, err = time.Parse(timeFormat, meeting.EndDate)

	if err != nil {
		return fmt.Errorf("EndTime format is illegal")
	}

	if meeting.StartDate > meeting.EndDate {
		return fmt.Errorf("StartTime is after EndTime")
	}

	meetings := entity.AllMeetings.FindBy(func(meeting *Meeting) bool {
		if meeting.Speecher == newMeeting.Speecher {
			return true
		}
		for _, participator := range meeting.Participators {
			if participator == newMeeting.Speecher {
				return true
			}
		}
		return false
	})

	for _, participatorName := range newMeeting.Participators {
		results = append(results, entity.MeetingModel.FindBy(func(m *meeting) bool {
			if m.Speecher == participatorName {
				return true
			}
			for _, participator := range m.Participators {
				if participator == participatorName {
					return true
				}
			}
			return false
		})...)
	}

	for _, meeting := range meetings {
		if endTime > meeting.StartTime && meeting.EndTime > startTime {
			return fmt.Errorf("Conflit with meeting %s", meeting.Title)
		}
	}

	return nil

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
	if entity.AllMeetings.FindByTitle(title).size != 0 {
		return fmt.Errorf("The name same as %s has been added", title)
	}

	speecherName := entity.CurrSession.GetCurUser()

	// check if the participators exist
	for _, partiName := range participatorName {
		if len(entity.AllUsers.FindByName(partiName)) == 0 {
			return fmt.Errorf("User %s doesn't exist", partiName)
		}
		if participatorName == speecherName {
			return fmt.Errorf("You can't be the participator")
		}
	}

	newMeeting := &meeting{
		Title:         title,
		Sponsor:       speecherName,
		Participators: participatorName,
		StartDate:     startTime,
		EndDate:       endTime,
	}

	if err := validateMeeting(newMeeting); err != nil {
		return err
	}

	entity.AllMeetings.AddMeeting(newMeeting)
	return nil
}

func DeleteFromMeeting(title string) error {

	// check if someone has logged in
	if !entity.CurrSession.HasLoggedIn() {
		return fmt.Errorf("You have not logged in")
	}

	meeting := entity.MeetingModel.FindByTitle(title)
	if meeting.Sponsor != entity.CurrSession.GetCurUser() {
		entity.MeetingModel.DeleteParticipatorFromMeeting(meeting, entity.CurrSession.GetCurUser())
		if len(entity.MeetingModel.FindByTitle(title).Participators) == 0 {
			//entity.MeetingModel.deleteMeeting(title)
		}
		return nil
	}
	//entity.MeetingModel.deleteMeeting(title)
	return nil
}

// delete the meeting whose name is title
func DeleteFromMeetingByTitle(title string) error {
	if meeting, err := entity.QueryMeeting(title); err == nil {
		entity.DeleteMeeting(title)
		return nil
	} else {
		return fmt.Errorf("no meeting to be deleted")
	}
}

//删除与会者
func DeleteParticipatorFromMeeting(title string, participatorNames []string) error {
	//登陆检查
	if !entity.CurrSession.HasLoggedIn() {
		return fmt.Errorf("No one has logged in")
	}
	//检查会议是否存在
	curMeeting := entity.MeetingModel.FindByTitle(title)
	if curMeeting == nil {
		return fmt.Errorf("Meeting %s doesn't exist", title)
	}
	//检查参与者是否存在
	for _, participatorName := range participatorNames {
		participator := entity.AllUsers.FindByName(participatorName)
		if participator == nil {
			return fmt.Errorf("User %s doesn't exist", participatorName)
		}

		flag := false

		for _, curParticipator := range curMeeting.Participators {
			if curParticipator == participatorName {
				entity.MeetingModel.DeleteParticipatorFromMeeting(curMeeting, curParticipator)
				flag = true
				// check if the participator is 0
				if len(entity.MeetingModel.FindByTitle(title).Participators) == 0 {
					//entity.MeetingModel.deleteMeeting(title)  BUG * 3
				}
				break
			}
		}
		if !flag {
			return fmt.Errorf("Participator %s is not in the meeting %s", participatorName, title)
		}
	}
	return nil
}

func AddParticipatorToMeeting(title string, participatorNames []string) error {

	if !entity.CurrSession.HasLoggedIn() {
		return fmt.Errorf("No one has logged in")
	}

	targetMeeting := entity.AllMeetings.FindByTitle(title)
	if targetMeeting == nil {
		return fmt.Errorf("Meeting %s doesn't exist", title)
	}

	for _, participatorName := range participatorNames {
		participator := entity.AllUsers.FindByName(participatorName)
		if participator == nil {
			return fmt.Errorf("User %s doesn't exist", participatorName)
		}

		meetings := entity.AllMeetings.FindBy(func(meeting *Meeting) bool {
			if meeting.Speecher == participatorName {
				return true
			}
			for _, pName := range meeting.Participators {
				if pName == participatorName {
					return true
				}
			}
			return false
		})

		for _, meeting := range meetings {
			if targetMeeting.EndDate > meeting.StartDate && meeting.EndDate > targetMeeting.StartDate {
				return fmt.Errorf("Conflit with meeting %s", meeting.Title)
			}
		}
	}

	for _, participatorName := range participatorNames {
		entity.AllMeetings.AddParticipatorToMeeting(curMeeting, participatorName)
	}

	return nil
}
