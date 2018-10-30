package service

import (
	"AgendaGo/entity"
	"fmt"
	"time"
)

type meeting = entity.Meeting

const timeFormat = "2018-10-28/13:08:22"

func validateNewMeeting(meeting *meeting) error {

	if err := validateNewTimeInterval(meeting.StartTime, meeting.EndTime); err != nil {
		return err
	}

	if flag, err := validateNewMeetingTime(meeting); !flag {
		return err
	}

	return nil
}

func validateNewTimeInterval(startTime string, endTime string) error {
	if len(startTime) == 0 {
		return fmt.Errorf("Empty start time")
	}
	_, err := time.Parse(timeFormat, startTime)

	if err != nil {
		return fmt.Errorf("Illegal StartTime format")
	}

	if len(endTime) == 0 {
		return fmt.Errorf("Empty end time")
	}
	_, err = time.Parse(timeFormat, endTime)

	if err != nil {
		return fmt.Errorf("Illegal EndTime format")
	}

	if startTime > endTime {
		return fmt.Errorf("StartTime is after the EndTime")
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
			return fmt.Errorf("Participator %s doesn't exist", participatorName)
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

//增加与会者
func AddParticipatorToMeeting(title string, participatorNames []string) error {
	//登陆检查
	if !entity.CurrSession.HasLoggedIn() {
		return fmt.Errorf("No one has logged in")
	}
	curMeeting := entity.MeetingModel.FindByTitle(title)
	if curMeeting == nil {
		return fmt.Errorf("Meeting %s doesn't exist", title)
	}

	for _, participatorName := range participatorNames {
		participator := entity.AllUsers.FindByName(participatorName)
		if participator == nil {
			return fmt.Errorf("Participator %s doesn't exist", participatorName)
		}
		//不处理
		// results := entity.MeetingModel.FindBy(func(m *meeting) bool {
		// 	// check from speecher field
		// 	if m.Sponsor == participatorName {
		// 		return true
		// 	}
		// 	// check from participator field
		// 	for _, participator := range m.Participators {
		// 		if participator == participatorName {
		// 			return true
		// 		}
		// 	}
		// 	return false
		// })

		// _, err := validateFreeTime(curMeeting.StartTime, curMeeting.EndTime, results)
		// if err != nil {
		// 	return err
		// }

	}

	for _, participatorName := range participatorNames {
		entity.MeetingModel.AddParticipatorToMeeting(curMeeting, participatorName)
	}

	return nil
}
