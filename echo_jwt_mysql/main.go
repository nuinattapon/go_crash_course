package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	sqlx "github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Define mysqlDB as a global variable
var mysqlDB *sqlx.DB
var jwtSecret = "ihcvRewKHN6=DL|J2ibaV+i&W"

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

	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time = ${time_rfc3339}, method=${method}, uri=${uri}, status=${status},latency=${latency_human}\n",
	}))

	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	h := &handler{}
	e.POST("/login", h.login)
	e.GET("/private", h.private, isLoggedIn)
	e.POST("/token", h.token, isLoggedIn) // refresh access token
	e.GET("/admin", h.private, isLoggedIn, isAdmin)

	// Start echo and handle errors
	// Start server
	port := 1323
	if err := e.Start(fmt.Sprintf(":%d", port)); err != nil {
		e.Logger.Fatal(err.Error())
	}
}
