package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Handler
func helloHandler(c echo.Context) error {
	return c.HTML(http.StatusOK, "<strong>Hello, 世界, สวัสดี</strong>")
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status},latency=${latency_human}\n",
	}))

	// e.Use(middleware.Gzip())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", helloHandler)

	// Start server
	port := 1323
	if err := e.Start(fmt.Sprintf(":%d", port)); err != nil {
		e.Logger.Fatal(err.Error())
	}

}
