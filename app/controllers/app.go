package controllers

import (
	"github.com/revel/revel"
	// alias "do_it/app/models/user"
	_ "errors"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-gorp/gorp"
	"log"
	"fmt"
	"time"
	_ "net/http"
	"strconv"
	"crypto/sha256"
	"encoding/hex"
	"os"
)

type task struct {
	Id int
	Title string 
	Status bool
}

type task_list struct {
	I_U 	[]task
	I_NU	[]task
	NI_U	[]task
	NI_NU	[]task
}

func return_tasks(username string) task_list {
	
	fmt.Println("Finding tasks for ", username)

	var tl task_list

	rows , err := dbmap.Query("select id, title, priority, status from task where username = ?", username)
	for rows.Next() {
		var id int
		var title string
		var priority int
		var status bool
		if err = rows.Scan(&id, &title, &priority, &status); err != nil {
			fmt.Println(err)
		}
		var t task
			t.Id 	= id
			t.Title = title
			t.Status = status

		if priority == 1 {
			tl.I_U = append(tl.I_U, t)
		} else if priority == 2 {
			tl.I_NU = append(tl.I_NU, t)
		} else if priority == 3 {
			tl.NI_U = append(tl.NI_U, t)
		} else if priority == 4 {
			tl.NI_NU = append(tl.NI_NU, t)
		}
	}
	return tl
}


type App struct {
	*revel.Controller
}

func (c App) Toggle() revel.Result {

	id 				:= c.Params.Route.Get("id")
	username 		:= c.Session["username"].(string)

	rows, err := dbmap.Query("select status from task where username = ? and id = ?", username, id)

	var old_status bool

	for rows.Next() {
		var status bool

		if err = rows.Scan(&status); err != nil {
			fmt.Println(err)
		}
		old_status = status
	}

	old_status = !old_status

	_, err = dbmap.Query("update task set status = ? where id = ?", old_status, id)

	return c.Redirect(App.Home)
}

func (c App) Delete() revel.Result {

	id 				:= c.Params.Route.Get("id")
	username 		:= c.Session["username"].(string)

	_, err := dbmap.Query("delete from task where username = ? and id = ?", username, id)

	if err != nil {
		log.Println(err)
	}
	
	return c.Redirect(App.Home)
}

func (c App) Add() revel.Result {
	title 			:= c.Params.Query.Get("title")
	priority, err 	:= strconv.Atoi(c.Params.Query.Get("priority"))
	username 		:= c.Session["username"].(string)
	t1 				:= newTask(title, username, priority)
	err = dbmap.Insert(&t1)
	if err != nil {
		fmt.Println(err)
	}
	return c.Redirect(App.Home)
}

func (c App) Home() revel.Result {
	// Allow bulk addition of task 

	username := c.Session["username"]
	// _ = username
	tasks := return_tasks(username.(string))
	
	fmt.Println(tasks.I_U)

    c.ViewArgs["username"] = c.Session["username"]
	// return c.RenderTemplate("app/home.html")
	I_U := tasks.I_U
	I_NU := tasks.I_NU
	NI_U := tasks.NI_U
	NI_NU := tasks.NI_NU

	return c.Render(I_U, I_NU, NI_U, NI_NU)
}

func (c App) Index() revel.Result {

	if c.Session["logged_in"] == "true" {
		c.ViewArgs["username"] = c.Session["username"]
		return c.Redirect(App.Home)
	} else {
		return c.RenderTemplate("app/index.html")
	}	
}

func (c App) Signup() revel.Result {

	username := c.Params.Get("username")
	password := c.Params.Get("password")

	count, _ := dbmap.SelectInt("select count(*) from user where username = ?;", username)
	if count == 0 {
			password = Hasher(password)
			_ , err := dbmap.Query("insert into user (username, password, created_at) values (?, ?, ?);", username, password, time.Now())
			if err != nil {
				log.Print(err)
			}
	} else {
		log.Print("Username already exists")
	}
	return c.Redirect(App.Index)
}

func (c App) Login() revel.Result {

	username := c.Params.Get("username")
	password := c.Params.Get("password")

	password = Hasher(password)

	rows, err := dbmap.Query("select exists (select * from user where username = ? and password = ?);", username, password)

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
			var a int64

			if err := rows.Scan(&a); err != nil {
				log.Fatal(err)
				fmt.Println(err)
			}

			if a == 1 {
				c.Session["logged_in"] = "true"
				c.Session["username"] = username
				c.ViewArgs["username"] = username

				return c.Redirect(App.Home)
			} 
		}
	return c.Redirect(App.Index)
}

func (c App) Logout() revel.Result {
	delete(c.Session, "logged_in")
	delete(c.Session, "username")
	return c.Redirect(App.Index)
}



// USER MODEL 
type User struct {
	Id			int64
	Username	string
	Password	string
	Created_at		time.Time
}

// TASK MODEL
type Task struct {
	Id			int64
	Title		string
	Priority	int
	Status		bool
	Username		string
	Created_at		time.Time
	Scheduled_for	time.Time
	Completed_at	time.Time
}

func newTask(title string, user string, priority int) Task {

	return Task{
		Title: 			title,
		Priority: 		priority,
		Status:			false,
		Username:		user,
		Created_at: 	time.Now(),
		Scheduled_for: 	time.Now(),
		Completed_at:	time.Now(),
	}
}

// Global database references
var db *sql.DB
var dbmap *gorp.DbMap

// Database settings
var db_name = os.Getenv("DB_NAME")
var db_user = os.Getenv("DB_USER")
var db_pw   = os.Getenv("DB_PASSWORD")

func InitDB() {

	var err error

	db, err = sql.Open("mysql", db_user + ":" + db_pw + "@tcp(127.0.0.1:3306)/" + db_name + "?parseTime=true")
	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	if err != nil {
		fmt.Println("Failed to connect to database: ")
		log.Panic(err)
	} else {
		err = db.Ping()

		if err != nil {
			fmt.Println("Failed to ping database: ")
			log.Panic(err)
		} else {
			fmt.Println("Database connected.")
		}
	}

 	_ = dbmap.AddTableWithName(User{}, "username").SetKeys(true, "Id")
 	_ = dbmap.AddTableWithName(Task{}, "task").SetKeys(true, "Id")
 	dbmap.CreateTablesIfNotExists()
}




func Hasher(value string) string {
	var salt = os.Getenv("ENCRYPTION_SALT")
	value = value + salt
	h := sha256.New()
	h.Write([]byte(value))
    sha256_hash := hex.EncodeToString(h.Sum(nil))
    return sha256_hash
}
