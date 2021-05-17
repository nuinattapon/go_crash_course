package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/godror/godror"
	"github.com/jmoiron/sqlx"
)

// Define db as a global variable
var db *sqlx.DB
var timezone *time.Location

/*
DROP TABLE nui.user2;
CREATE TABLE nui.user2 (
    id INT GENERATED ALWAYS AS IDENTITY (START WITH 1 INCREMENT BY 1) not null primary key,
    user_name VARCHAR2(255),
    email VARCHAR2(255),
    hashed_password VARCHAR2(255),
    is_admin INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP );

*/

type User struct {
	ID             int64     `db:"ID" json:"id"`
	UserName       string    `db:"USER_NAME" json:"user_name"`
	Email          string    `db:"EMAIL" json:"email"`
	HashedPassword string    `db:"HASHED_PASSWORD" json:"-"`
	IsAdmin        bool      `db:"IS_ADMIN" json:"is_admin"`
	CreatedAt      time.Time `db:"CREATED_AT" json:"created_at"`
	UpdatedAt      time.Time `db:"UPDATED_AT" json:"updated_at"`
}

func printSqlResult(result sql.Result) {
	lastInsertId, _ := result.LastInsertId()
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Execution Result: lastInsertId %d - rowsAffected %d\n", lastInsertId, rowsAffected)
}
func printUsers(users []User) {
	fmt.Printf("%-2s %-5s %-9s %-21s %-21s\n", "--", "----", "--------", "----------", "----------")
	fmt.Printf("%-2s %-5s %-9s %-21s %-21s\n", "ID", "Name", "Is admin", "Created At", "Updated At")
	fmt.Printf("%-2s %-5s %-9s %-21s %-21s\n", "--", "----", "--------", "----------", "----------")
	for _, u := range users {
		printUser(u)
	}
}

func printUser(user User) {
	adminStr := map[bool]string{true: "admin", false: "not admin"}[user.IsAdmin]
	fmt.Printf("%2d %-5s %-9s %-21s %-21s\n", user.ID, user.UserName, adminStr, user.CreatedAt.In(timezone).Format(time.RFC822Z), user.UpdatedAt.In(timezone).Format(time.RFC822Z))
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
	db, err = sqlx.Open("godror", "nui/Welcome1Welcome1@atp_high")
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	} else {
		log.Println("Database connection is initialized")
	}
	// db.SetMaxOpenConns(20)
}

func main() {
	var err error
	defer db.Close()

	// Check if we can query from oracle dual table
	rows, err := db.Query("select sysdate from dual")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	var thedate time.Time
	for rows.Next() {

		rows.Scan(&thedate)
	}
	rows.Close()
	fmt.Printf("\nCurrent time in ATP is %s\n\n", thedate.In(timezone).Format(time.RFC822Z))

	// Select is used to query multiple rows
	users := []User{}
	err = db.Select(&users, "SELECT * FROM nui.user2 ORDER by id FETCH FIRST 100 ROWS ONLY")
	if err != nil && err != sql.ErrNoRows {
		panic(err.Error()) // proper error handling instead of panic in your app
	} else if err == sql.ErrNoRows || len(users) == 0 {
		fmt.Println("Data before runing transaction")
		fmt.Println("No rows in the table!")
	} else {
		fmt.Println("Data before runing transaction")
		printUsers(users)
	}

	// Get is used for query a single row
	user_name := "test"
	user := User{}
	err = db.Get(&user, "SELECT * FROM nui.user2 WHERE user_name = &1 ORDER by id FETCH FIRST 100 ROWS ONLY", user_name)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("\nCan not find a user with user_name '%s'\n", user_name)
		} else {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	} else {
		fmt.Printf("\nA user with user_name '%s' is found\n", user_name)
		printUser(user)
	}

	updated_at := time.Now()

	tx := db.MustBegin()

	result, err := tx.Exec("UPDATE nui.user2 set updated_at = &1", updated_at)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("\nUPDATE nui.user2 set updated_at = &1")
	printSqlResult(result)

	// result, err = tx.Exec("DELETE from nui.user2 where user_name = 'test'")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// printSqlResult(result)
	result, err = tx.Exec(`
	INSERT into nui.user2  
	(user_name, email, hashed_password,is_admin,created_at,updated_at) 
	values (&1,&2,&3,&4,&5,&6)`, "test", "test@test.com", "$2a$11$fTDn/IzGVYpj5C2P1QewPOZwxuVpsHjH3go0YORqfDDk/G7UhZWna", 0, updated_at, updated_at)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("\nINSERT into nui.user2 (user_name, email, hashed_password,is_admin,created_at,updated_at) ...")
	printSqlResult(result)
	// tx.Rollback()
	tx.Commit()

	// Select
	users = []User{}
	err = db.Select(&users, "SELECT * FROM nui.user2 ORDER by id FETCH FIRST 100 ROWS ONLY")
	if err != nil && err != sql.ErrNoRows {
		panic(err.Error()) // proper error handling instead of panic in your app
	} else if err == sql.ErrNoRows || len(users) == 0 {
		fmt.Println("\nData after runing transaction")
		fmt.Println("No rows in the table!")
	} else {
		fmt.Println("\nData after runing transaction")
		printUsers(users)
	}

}
