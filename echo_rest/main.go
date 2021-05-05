package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	sqlx "github.com/jmoiron/sqlx"
	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
)

// Define mysqlDB as a global variable
var mysqlDB *sqlx.DB

// In addition to echo request handlers
// using a special context including
// all kinds of utilities, generated errors
// can be returned to handle them easily

func PingGetHandler(e echo.Context) error {
	return e.String(http.StatusOK, "pong")
}

func VersionGetHandler(e echo.Context) error {
	return e.String(http.StatusOK, "v1.0")
}

func HelloGetHandler(e echo.Context) error {
	return e.HTML(http.StatusOK, "<h1>Hello</h1>")
}

// This simple struct will be deserialized
// and processed in the request handler
type Test struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type TestJSON struct {
	ID   int    `db:"id" json:"id,omitempty"`
	Name string `db:"name" json:"name,omitempty"`
	Data string `db:"data" json:"data"`
}

func TestJSONGetHandler(c echo.Context) error {
	// Execute the query
	// We use sqlx syntax here in stead of golang sql
	testSlice := []TestJSON{}
	err := mysqlDB.Select(&testSlice, "SELECT id, name, data FROM acme.test_json2 LIMIT 100")

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	return c.JSON(http.StatusOK, testSlice)
}

func TestGetHandler(c echo.Context) error {
	// Execute the query
	// We use sqlx syntax here in stead of golang sql
	testSlice := []Test{}
	err := mysqlDB.Select(&testSlice, "SELECT id, name FROM acme.test LIMIT 100")

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// In this case we can return the JSON
	// function with our body as errors
	// thrown by this will be handled
	return c.JSON(http.StatusOK, testSlice)
}

func TestGetHandler2(c echo.Context) error {
	// Create response object
	// fmt.Println(e.ParamNames())
	// fmt.Println(e.ParamValues())
	// to get query string parameters
	// - e.Request.URL.Query().Get("bar")

	// We use sqlx syntax here in stead of golang sql
	fmt.Println(c.Param("id"))
	testSlice := []Test{}
	err := mysqlDB.Select(&testSlice, "SELECT id, name FROM acme.test WHERE id = ?", c.Param("id"))

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	if len(testSlice) == 1 {
		return c.JSON(http.StatusOK, testSlice[0])

	} else if len(testSlice) == 0 {
		return c.JSON(http.StatusNotFound, testSlice)
	} else {
		return c.JSON(http.StatusOK, testSlice)
	}
}

func TestPostHandler(c echo.Context) error {
	// Similar to the gin implementation,
	// we start off by creating an
	// empty request body struct
	test := &Test{}
	if err := c.Bind(test); err != nil {
		return err
	}
	// Execute the query
	testSlice := []Test{}

	if err := mysqlDB.Select(&testSlice, "SELECT id, name FROM acme.test WHERE name = ?", test.Name); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	if len(testSlice) != 0 {
		test.ID = -1
		return c.JSON(http.StatusMethodNotAllowed, test)
	}
	tx := mysqlDB.MustBegin()

	if result, err := tx.Exec("INSERT INTO acme.test (name) VALUES( ? )", test.Name); err != nil {
		panic(err.Error())
	} else {
		tx.Commit()
		lastInsertID, _ := result.LastInsertId()
		test.ID = int(lastInsertID)
		return c.JSON(http.StatusOK, test)
	}
}

type User struct {
	ID           int          `db:"id" json:"id"`
	FirstName    string       `db:"first_name" json:"first_name"`
	LastName     string       `db:"last_name" json:"last_name"`
	Email        string       `db:"email" json:"email"`
	Password     string       `db:"password" json:"password,omitempty"`
	Location     string       `db:"location" json:"location"`
	Department   string       `db:"dept" json:"department"`
	IsAdmin      int          `db:"is_admin" json:"is_admin"`
	RegisterDate sql.NullTime `db:"register_date" json:"register_date"`
	Age          int          `db:"age" json:"age"`
}

func UserGetHandler(c echo.Context) error {
	// Execute the query
	// We use sqlx syntax here in stead of golang sql
	userSlice := []User{}
	err := mysqlDB.Select(&userSlice, "SELECT * FROM acme.users LIMIT 100")

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// In this case we can return the JSON
	// function with our body as errors
	// thrown by this will be handled
	return c.JSON(http.StatusOK, userSlice)
}

func main() {
	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called "mysql"
	var err error

	mysqlDB, err = sqlx.Open("mysql", "nattapon:Welcome1@tcp(192.168.1.6:3306)/acme?parseTime=true")
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	mysqlDB.SetMaxOpenConns(20)
	mysqlDB.SetMaxIdleConns(10)
	// mysqlDB.SetConnMaxLifetime(time.Duration(3600)*time.Second)

	// defer the close till after the main function has finished
	// executing
	defer mysqlDB.Close()
	log.Println("Successfully connecting to MySQL Database!")

	// Create echo instance
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status},latency=${latency_human}\n",
	}))
	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())

	// Add endpoint route for /ping /version and /hello
	e.GET("/ping", PingGetHandler)
	e.GET("/version", VersionGetHandler)
	e.GET("/hello", HelloGetHandler)

	// Add endpoint route for /test
	e.GET("/test", TestGetHandler)
	e.GET("/test/:id", TestGetHandler2)
	e.POST("/test", TestPostHandler)

	// Add endpoint route for /test_json
	e.GET("/test_json", TestJSONGetHandler)

	// Add endpoint route for /user
	e.GET("/user", UserGetHandler)

	// Start echo and handle errors
	// Start server
	port := 8002
	if err := e.Start(fmt.Sprintf(":%d", port)); err != nil {
		e.Logger.Fatal(err.Error())
	}

	// port := 8002
	// if err := e.StartTLS(fmt.Sprintf(":%d", port), "./certs/tls-cert.crt", "./certs/tls-key.key"); err != nil {
	// 	e.Logger.Fatal(err.Error())
	// }
}
