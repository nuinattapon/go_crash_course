package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
)

// Define mysqlDB as a global variable
var mysqlDB *sql.DB

func init() {
	log.Printf("GOMAXPROCS: %+v, NumCPU: %+v\n",
		runtime.GOMAXPROCS(-1), runtime.NumCPU())

}

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
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func UserGetHandler(e echo.Context) error {
	// Execute the query
	results, err := mysqlDB.Query("SELECT id, name FROM acme.test LIMIT 100")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	testSlice := []Test{}

	for results.Next() {
		var test Test
		// for each row, scan the result into our tag composite object
		err = results.Scan(&test.ID, &test.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		testSlice = append(testSlice, test)
	}

	// stat := db.Stats()
	// fmt.Printf("%+v", stat)

	// In this case we can return the JSON
	// function with our body as errors
	// thrown by this will be handled
	return e.JSON(http.StatusOK, testSlice)
}
func UserGetHandler2(e echo.Context) error {
	// Create response object
	// fmt.Println(e.ParamNames())
	// fmt.Println(e.ParamValues())
	// to get query string parameters
	// - e.Request.URL.Query().Get("bar")
	fmt.Println(e.Param("id"))
	// Execute the query
	results, err := mysqlDB.Query("SELECT id, name FROM acme.test WHERE id = ?", e.Param("id"))
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	testSlice := []Test{}

	for results.Next() {
		var test Test
		// for each row, scan the result into our tag composite object
		err = results.Scan(&test.ID, &test.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		testSlice = append(testSlice, test)
	}

	// stat := db.Stats()
	// fmt.Printf("%+v", stat)

	// In this case we can return the JSON
	// function with our body as errors
	// thrown by this will be handled
	if len(testSlice) == 1 {
		return e.JSON(http.StatusOK, testSlice[0])

	} else if len(testSlice) == 0 {
		return e.JSON(http.StatusNotFound, testSlice)
	} else {
		return e.JSON(http.StatusOK, testSlice)
	}
}

func UserPostHandler(e echo.Context) error {
	// Similar to the gin implementation,
	// we start off by creating an
	// empty request body struct
	test := &Test{}
	err := e.Bind(test)
	if err != nil {

		return err
	}
	// Execute the query
	results, err := mysqlDB.Query("SELECT id, name FROM acme.test WHERE name = ?", test.Name)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	testSlice := []Test{}

	for results.Next() {
		var test Test
		// for each row, scan the result into our tag composite object
		err = results.Scan(&test.ID, &test.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		testSlice = append(testSlice, test)
	}

	if len(testSlice) != 0 {
		test.ID = 0
		return e.JSON(http.StatusMethodNotAllowed, test)
	}

	// Insert a name into acme.test table
	insertStmt, err := mysqlDB.Prepare("INSERT INTO acme.test (name) VALUES( ? )")
	if err != nil {
		panic(err.Error())
	}
	defer insertStmt.Close()

	_, err = insertStmt.Exec(test.Name)
	if err != nil {
		panic(err.Error())
	}

	// Get id from acme.test
	// Execute the query
	// test2 := &Test{}

	var insertID int64
	err = mysqlDB.QueryRow("select LAST_INSERT_ID()").Scan(&insertID)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	test.ID = int(insertID)
	return e.JSON(http.StatusOK, test)
}

func main() {
	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called "mysql"
	var err error

	mysqlDB, err = sql.Open("mysql", "nattapon:Welcome1@tcp(192.168.1.6:3306)/mysql")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	mysqlDB.SetMaxOpenConns(10)

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
	e.GET("/test", UserGetHandler)
	e.GET("/test/:id", UserGetHandler2)
	e.POST("/test", UserPostHandler)
	// Add endpoint route for /test/<username>
	// Start echo and handle errors
	// Start server
	port := 8002
	if err := e.Start(fmt.Sprintf(":%d", port)); err != nil {
		e.Logger.Fatal(err.Error())
	}
}
