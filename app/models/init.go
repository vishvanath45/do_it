package models

import (
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/go-gorp/gorp"
	_ "log"
	_ "fmt"
)

// // Global database references
// var db *sql.DB
// var dbmap *gorp.DbMap

// // Database settings
// var db_name = "do_it"
// var db_user = "do_it_admin"
// var db_pw   = "do_it_password"

// func InitDB() {
// 	var err error

// 	db, err = sql.Open("mysql", db_user + ":" + db_pw + "@tcp(127.0.0.1:3306)/" + db_name)
// 	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

// 	if err != nil {
// 		fmt.Println("Failed to connect to database: ")
// 		log.Panic(err)
// 	} else {
// 		err = db.Ping()

// 		if err != nil {
// 			fmt.Println("Failed to ping database: ")
// 			log.Panic(err)
// 		} else {
// 			fmt.Println("Database connected.")
// 		}
// 	}

//  	_ = dbmap.AddTableWithName(User{}, "user").SetKeys(true, "Id")
//  	_ = dbmap.AddTableWithName(Task{}, "task").SetKeys(true, "Id")
//  	dbmap.CreateTablesIfNotExists()
// }
