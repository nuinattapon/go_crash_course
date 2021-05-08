package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Define db as a global variable
var db *sqlx.DB
var timezone *time.Location

type User struct {
	ID             int64     `db:"uid" json:"id"`
	UserName       string    `db:"user_name" json:"user_name"`
	Email          string    `db:"email" json:"email"`
	HashedPassword string    `db:"hashed_password" json:"-"`
	IsAdmin        bool      `db:"is_admin" json:"is_admin"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}

func init() {
	var err error
	// Initialize timezone variable
	timezone, err = time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err.Error())
	} else {
		log.Println("Timezone is initialzied to 'Asia/Bangkok' ")

	}

	// Initialize db variable
	db, err = sqlx.Open("mysql", "dev:Welcome1@tcp(192.168.1.6:3306)/nui?parseTime=true")
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	} else {
		log.Println("Database connection is initialized")
	}
	db.SetMaxOpenConns(20)
}

func main() {

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called "mysql"
	var err error

	// Select is used to query multiple rows
	users := []User{}
	err = db.Select(&users, "SELECT * FROM nui.user ORDER by uid LIMIT 100")
	if err != nil && err != sql.ErrNoRows {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Println("Data before runing transaction")

	printUsers(users)

	// Get is used for query a single row
	user_name := "test"
	user := User{}
	err = db.Get(&user, "SELECT * FROM nui.user WHERE user_name = ? LIMIT 100", user_name)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("\nCan not find a user with user_name '%s'\n", user_name)
		} else {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

	}
	fmt.Printf("\nA user with user_name '%s' is found\n", user_name)
	printUser(user)

	updated_at := time.Now()

	tx := db.MustBegin()

	result, err := tx.Exec("UPDATE nui.user set updated_at = ?", updated_at)
	if err != nil {
		panic(err.Error())
	}
	printSqlResult(result)

	result, err = tx.Exec("DELETE from nui.user where user_name = 'test'")
	if err != nil {
		panic(err.Error())
	}
	printSqlResult(result)

	result, err = tx.Exec(`
	INSERT into nui.user  
	(user_name, email, hashed_password,is_admin,created_at,updated_at) 
	values (?,?,?,?,?,?)`, "test", "test@test.com", "$2a$11$fTDn/IzGVYpj5C2P1QewPOZwxuVpsHjH3go0YORqfDDk/G7UhZWna", 0, updated_at, updated_at)
	if err != nil {
		panic(err.Error())
	}
	printSqlResult(result)

	// tx.Rollback()
	tx.Commit()
	fmt.Println("\nData after runing transaction")

	// Select
	users = []User{}
	err = db.Select(&users, "SELECT * FROM nui.user ORDER BY uid LIMIT 100")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	printUsers(users)

}

func printSqlResult(result sql.Result) {
	lastInsertId, _ := result.LastInsertId()
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Execution Result: lastInsertId %d - rowsAffected %d\n", lastInsertId, rowsAffected)
}
func printUsers(users []User) {
	for _, u := range users {
		printUser(u)
	}
}

func printUser(user User) {
	adminStr := map[bool]string{true: "admin", false: "not admin"}[user.IsAdmin]
	fmt.Printf("%2d - %-5s - %-9s - %s - %s\n", user.ID, user.UserName, adminStr, user.CreatedAt.In(timezone).Format(time.RFC822Z), user.UpdatedAt.In(timezone).Format(time.RFC822Z))
}
