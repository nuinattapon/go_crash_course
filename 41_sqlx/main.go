package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Define mysqlDB as a global variable
var mysqlDB *sqlx.DB

type User struct {
	ID             int64     `db:"uid" json:"id"`
	UserName       string    `db:"user_name" json:"user_name"`
	Email          string    `db:"email" json:"email"`
	HashedPassword string    `db:"hashed_password" json:"-"`
	IsAdmin        bool      `db:"is_admin" json:"is_admin"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}

func main() {

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called "mysql"
	var err error

	mysqlDB, err = sqlx.Open("mysql", "nui:Welcome1@tcp(192.168.1.6:3306)/nui?parseTime=true")
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	log.Println("Database connection is initialized")
	mysqlDB.SetMaxOpenConns(20)

	// Select is used to query multiple rows
	userSlice := []User{}
	err = mysqlDB.Select(&userSlice, "SELECT * FROM nui.user ORDER by uid LIMIT 100")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for i, u := range userSlice {
		var adminStr string
		if u.IsAdmin {
			adminStr = "admin"
		} else {
			adminStr = "not admin"
		}
		fmt.Printf("%d - %-5s %-9s %v %v\n", i, u.UserName, adminStr, u.CreatedAt, u.UpdatedAt)
	}

	// Get is used for query a single row
	user_name := "nui"

	user := User{}
	err = mysqlDB.Get(&user, "SELECT * FROM nui.user WHERE user_name = ? LIMIT 100", user_name)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("%+v\n", user.UpdatedAt)

	// Exec
	// var tx sqlx.Tx

	tx := mysqlDB.MustBegin()

	updated_at := time.Now()
	tx.MustExec("UPDATE nui.user set updated_at = ?", updated_at)

	// tx.Rollback()
	tx.Commit()

	// Exec
	// var tx sqlx.Tx

	tx = mysqlDB.MustBegin()

	updated_at = time.Now()
	// tx.MustExec("DELETE from nui.user where user_name = 'test'")
	tx.MustExec(`
	INSERT into nui.user 
	(user_name, email, hashed_password,is_admin,created_at,updated_at) 
	values (?,?,?,?,?,?)`,
		"test", "test@test.com", "$2a$11$fTDn/IzGVYpj5C2P1QewPOZwxuVpsHjH3go0YORqfDDk/G7UhZWna", 0, updated_at, updated_at)

	// tx.Rollback()
	tx.Commit()
	// Select
	userSlice = []User{}
	err = mysqlDB.Select(&userSlice, "SELECT * FROM nui.user ORDER BY uid LIMIT 100")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for i, u := range userSlice {

		adminStr := map[bool]string{true: "admin", false: "not admin"}[u.IsAdmin]

		fmt.Printf("%d - %-5s %-9s %v %v\n", i, u.UserName, adminStr, u.CreatedAt, u.UpdatedAt)
	}
}
