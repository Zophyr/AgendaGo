package entity

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type UserDB struct {
	Data []User `json:"data"`
}

type Users struct {
	storage
	users map[string]*User
}

var AllUsers Users

func (allUsers *Users) AddUser(user *User) {
	defer allUsers.dump()
	allUsers.users[user.Username] = user
}

func (allUsers *Users) DeleteUser(user *User) {
	defer allUsers.dump()
	delete(allUsers.users, user.Username)
}

// use a filter to find appropriate users
func (allUsers *Users) FindBy(cond func(*User) bool) []User {
	result := []User{}
	for _, user := range allUsers.users {
		if cond(user) {
			result = append(result, *user)
		}
	}
	return result
}

func (allUsers *Users) FindByName(username string) []User {
	return allUsers.FindBy(func(user *User) bool {
		return username == user.Username
	})
}

func (allUsers *Users) IsMatchNamePass(username, password string) bool {
	result := allUsers.FindByName(username)
	return result[0].Password == password
}

func (allusers *Users) load() {
	var userDB UserDB
	allusers.storage.load(&userDB)
	for index, user := range userDB.Data {
		allusers.users[user.Username] = &userDB.Data[index]
	}
}

func (allusers *Users) dump() {
	var userDB UserDB
	for _, user := range allusers.users {
		userDB.Data = append(userDB.Data, *user)
	}
	allusers.storage.dump(&userDB)
}

func (allusers *Users) Init(path string) { // user call this function in the root cmd
	allusers.storage.path = "./data/user.json"
	allusers.users = make(map[string]*User)
	allusers.load()
}

func init() {
	addModel(&AllUsers, "user_data")
}
