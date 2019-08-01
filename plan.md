// this is safe
db.Query("SELECT name FROM users WHERE age=?", req.FormValue("age"))
// this allows sql injection.
db.Query("SELECT name FROM users WHERE age=" + req.FormValue("age"))


fmt.Println("Hello\nHello\nHello\n")


M0    Target audience - > Me

1. User Login   (DONE)
2. User should be able to create task, assign priority (DB)   (DONE)
3. Visualization (2x2 Matrix)
-- all tasks that have given scheduled_at
4. Task Moveover
-- Move to today (s_at)
-- Backlog maintain(All uncompleted task, will have option to move today)
5. Task points.


Task strikethrough



M1 

User stats
tasks (is_private)
Post to reddit on your behalf on subreddit r/getdisciplined @ given time.

M2 

Add Friends
notification system 
friends feed
(db maintain, state maintain)

---> Public Posts
M3 

add task status as -> ongoing
Integrate Pomodoro timer 




DB models
User 

type User struct {
	Id			int64
	Username	string
	Password	string
	Created		int64
}

type Task struct {
	Id			int64
	Title		string
	Priority	enum
	Status		bool
	UserId		int64
	Created_at		int64/dt
Scheduled_for int64/dt
    Completed_at  int64/dt 	
}


      {{range $I_U := .}}
        <u>{{$.status}}</u>
        <i>{{$.title}}</i>
      {{end}}







