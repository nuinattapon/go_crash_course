package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
)

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

func fib(n int64) int64 {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func FiboGetHandler2(c echo.Context) error {
	// Create response object
	// fmt.Println(c.ParamNames())
	// fmt.Println(c.ParamValues())
	// to get query string parameters
	// - c.Request.URL.Query().Get("bar")
	var n int64
	var err error

	if n, err = strconv.ParseInt(c.Param("id"), 10, 64); err != nil {
		n = 0
	}
	if n < 0 || n > 40 {
		return c.String(http.StatusNotAcceptable, "n should be between 0 and 40")
	}
	f := fib(n)
	// fmt.Println(f)
	returnString := fmt.Sprintf("%d", f)
	return c.String(http.StatusOK, returnString)

}

func main() {
	// Create echo instance
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
	}))
	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())

	// Add endpoint route for /ping /version and /hello
	e.GET("/ping", PingGetHandler)
	e.GET("/version", VersionGetHandler)
	e.GET("/hello", HelloGetHandler)

	// Add endpoint route for /test
	e.GET("/fibo/:id", FiboGetHandler2)

	// Start echo and handle errors
	// Start server
	port := 8002
	if err := e.Start(fmt.Sprintf(":%d", port)); err != nil {
		e.Logger.Fatal(err.Error())
	}
}
