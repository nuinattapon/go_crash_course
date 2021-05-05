package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type handler struct{}

// Most of the code is taken from the echo guide
// https://echo.labstack.com/cookbook/jwt
func (h *handler) private(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	var adminStatus string
	if isAdmin := claims["admin"].(bool); isAdmin {
		adminStatus = "admin"
	} else {
		adminStatus = "no admin"
	}
	// exp := claims["name"].(string)

	reply := fmt.Sprintf("Status %3d - Welcome %s.\nYou have %s privilege!", http.StatusOK, name, adminStatus)
	return c.String(http.StatusOK, reply)
}

// Most of the code is taken from the echo guide
// https://echo.labstack.com/cookbook/jwt
func (h *handler) login(c echo.Context) error {
	user_name := c.FormValue("username")
	password := c.FormValue("password")

	// Get user from mysql
	userSlice := []User{}
	err := mysqlDB.Select(&userSlice, "SELECT * FROM nui.user WHERE user_name = ? LIMIT 1", user_name)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	if len(userSlice) != 1 {
		return echo.ErrUnauthorized
	}
	user := userSlice[0]
	log.Printf("%+v\n", user)
	// Comparing the password with the hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password)); err == nil {
		// Create token
		log.Println("Password is correct!")
		token := jwt.New(jwt.SigningMethodHS512)
		// Set claims
		// This is the information which frontend can use
		// The backend can also decode the token and get admin etc.
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = user.UserName
		claims["admin"] = user.IsAdmin
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		// Generate encoded token and send it as response.
		// The signing string should be secret (a generated UUID          works too)
		t, err := token.SignedString([]byte(jwtSecret))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}
	return echo.ErrUnauthorized
}
