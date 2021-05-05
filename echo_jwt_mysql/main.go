package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status},latency=${latency_human}\n",
	}))
	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	h := &handler{}
	e.POST("/login", h.login)
	e.GET("/private", h.private, isLoggedIn)
	e.GET("/admin", h.private, isLoggedIn, isAdmin)

	// Start echo and handle errors
	// Start server
	port := 1323
	if err := e.Start(fmt.Sprintf(":%d", port)); err != nil {
		e.Logger.Fatal(err.Error())
	}
}
