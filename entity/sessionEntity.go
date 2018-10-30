package entity

type Session struct {
	CurrUser *User
}

var CurrSession Session

// return whether someone has been logged in or not
func (currSession *Session) HasLoggedIn() bool {
	return currSession.CurrUser != nil
}

// return the name of current user
func (currSession *Session) GetCurUserName() string {
	return currSession.CurrUser.Username
}
