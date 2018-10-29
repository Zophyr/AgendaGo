package entity

type Session struct {
	CurrUser *User
}

var CurrSession Session

func (currSession *Session) HasLoggedIn() bool {
	return currSession.CurrUser != nil
}
