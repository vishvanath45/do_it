package models

import (
	//"log"
	"time"
)

type Task struct {
	Id			int64
	Title		string
	Priority	string
	Status		bool
	UserId		int64
	Created		int64 	
}	

func newTask(title, priority string, user int64) Task {
	return Task{
		Title: 		title,
		Priority: 	priority,
		Status:		false,
		UserId:		user,
		Created: 	time.Now().UnixNano()/ int64(time.Millisecond),
	}
}