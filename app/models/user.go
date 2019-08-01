package models

import (
	// "log"
	"time"
)

type User struct {
	Id			int64
	Username	string
	Password	string
	Created		int64
}

func newUser(username, password string) User {
	return User{
		Username:	username,
		Password:	password,
		Created:	time.Now().UnixNano()/ int64(time.Millisecond),
	}
}
// func Finduser(number int) (bool, *User) {

// 	obj, err := dbmap.Get(User{}, number)
//  	emp := obj.(*User)

// 	if err != nil {
// 		log.Print("ERROR findEmployee: ")
// 	 	log.Println(err)
// 	}

// 	return (err == nil), emp
// }