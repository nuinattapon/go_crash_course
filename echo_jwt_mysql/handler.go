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

	// for k, v := range claims {
	// 	log.Printf("key = %v, type = %T, value = %v\n", k, v, v)
	// }
	id := claims["id"]
	name := claims["name"].(string)
	exp := time.Unix(int64(claims["exp"].(float64)), 0)

	var isAdminStr string
	if isAdmin := claims["admin"].(bool); isAdmin {
		isAdminStr = "admin"
	} else {
		isAdminStr = "no admin"
	}
	// exp := claims["name"].(string)

	reply1 := fmt.Sprintf("Status %3d - Welcome UID# %.0f (%s)\nYou have %s privilege!\n", http.StatusOK, id, name, isAdminStr)
	reply2 := fmt.Sprintf("Token expiry date: %s\n", exp.Format(time.RFC1123Z))
	reply3 := fmt.Sprintf("              Now: %s\n", time.Now().Format(time.RFC1123Z))
	var reply4 string
	if exp.Before(time.Now()) {
		reply4 = "Access is expired!"
	} else {
		reply4 = "Access is still valid."
	}
	reply := reply1 + reply2 + reply3 + reply4
	return c.HTML(http.StatusOK, reply)
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
		newTokenPair, err := generateTokenPair(user)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, newTokenPair)
	}
	return echo.ErrUnauthorized
}

type tokenReqBody struct {
	RefreshToken string `json:"refresh_token"`
}

// This is the api to refresh tokens
// Most of the code is taken from the jwt-go package's sample codes
// https://godoc.org/github.com/dgrijalva/jwt-go#example-Parse--Hmac
func (h *handler) token(c echo.Context) error {

	tokenReq := tokenReqBody{}
	c.Bind(&tokenReq)

	// Parse takes the token string and a function for looking up the key.
	// The latter is especially useful if you use multiple keys for your application.
	// The standard is to use 'kid' in the head of the token to identify
	// which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenReq.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(jwtSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Get the user record from database or
		// run through your business logic to verify if the user can log in
		uid := int64(claims["sub"].(float64))

		userSlice := []User{}
		err := mysqlDB.Select(&userSlice, "SELECT * FROM nui.user WHERE uid = ? LIMIT 1", uid)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		if len(userSlice) == 1 { //if we found a user from the uid
			user := userSlice[0]
			// log.Printf("%+v\n", user)
			newTokenPair, err := generateTokenPair(user)
			if err != nil {
				return err
			}

			return c.JSON(http.StatusOK, newTokenPair)
		}

		return echo.ErrUnauthorized
	}

	return err
}
