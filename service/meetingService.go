package service

import (
	"AgendaGo/entity"
	"fmt"
	"os"
)

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

func (model *meetingModel) deleteMeeting(title string) bool {
	if _, ok := model.meetings[title]; ok {
		delete(model.meetings, title)
		return true
	} else {
		fmt.Println(os.Stderr, "error:%s\n", "no such meeting to delete")
		return false
	}
}

type meeting = entity.Meeting

const timeFormat = "2018-10-28/13:08:22"

//增加会议

func AddMeeting(title string, participatorName []string, startTime string, endTime string) (err error) {
	// 没有登陆则返回报错
	if !entity.CurrSession.HasLoggedIn() {
		return fmt.Errorf("You have not logged in")
	}

	speecherName := entity.CurrSession.GetCurUser()
	newMeeting := &meeting{
		Title:         title,
		Sponsor:       speecherName,
		Participators: participatorName,
		StartDate:     startTime,
		EndDate:       endTime,
	}
	//验证
	// err = validateNewMeeting(newMeeting)
	// if err != nil {
	// 	return
	// }
	entity.MeetingModel.AddMeeting(newMeeting)
	return
}

func DeleteFromMeeting(title string) error {
	// 没有登陆则返回报错
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
