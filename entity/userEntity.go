package entity

type User struct {
	Username string
	Password string
	Email    string
	Phone    string
}

type Users struct {
	users map[string]*User
}

var AllUsers Users

func (allUsers *Users) AddUser(user *User) {
	allUsers.users[user.Username] = user
}

func (allUsers *Users) DeleteUser(user *User) {
	delete(allUsers.users, user.Username)
}

func (allUsers *Users) FindBy(cond func(*User) bool) []User {
	result := []User{}
	for _, user := range allUsers.users {
		if cond(user) {
			result = append(result, *user)
		}
	}
	return result
}
