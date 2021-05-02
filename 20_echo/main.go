package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// This is the response struct that will be
// serialized and sent back
type StatusResponse struct {
	Message string `json:"message"`
	User    string `json:"user"`
	Status  string `json:"status"`
}

// In addition to echo request handlers
// using a special context including
// all kinds of utilities, generated errors
// can be returned to handle them easily
func UserGetHandler(e echo.Context) error {
	// Create response object
	fmt.Println(e.ParamNames())
	fmt.Println(e.ParamValues())
	body := &StatusResponse{
		Message: "Hello, 世界, สวัสดี!",
		User:    e.Param("user"),
		Status:  "OK",
	}

	// In this case we can return the JSON
	// function with our body as errors
	// thrown by this will be handled
	return e.JSON(http.StatusOK, body)
}

// This simple struct will be deserialized
// and processed in the request handler
type RequestBody struct {
	Name string `json:"name"`
}

func UserPostHandler(e echo.Context) error {
	// Similar to the gin implementation,
	// we start off by creating an
	// empty request body struct
	requestBody := &RequestBody{}

	// Bind body to the request body
	// struct and check for potential
	// errors
	err := e.Bind(requestBody)
	if err != nil {
		// If an error was created by the
		// Bind operation, we can utilize
		// echo's request handler structure
		// and simply return the error so
		// it gets handled accordingly
		return err
	}

	body := &StatusResponse{
		Message: "Hello world from echo!",
		User:    requestBody.Name,
		Status:  "OK",
	}

	return e.JSON(http.StatusOK, body)
}

func main() {
	// Create echo instance
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status},latency=${latency_human}\n",
	}))

	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())

	// Add endpoint route for /users/<username>
	e.GET("/users/:user", UserGetHandler)

	// Add endpoint route for /users
	e.POST("/users", UserPostHandler)

	// Start echo and handle errors
	// Start server
	port := 8002
	if err := e.Start(fmt.Sprintf(":%d", port)); err != nil {
		e.Logger.Fatal(err.Error())
	}
}
