package main

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

func generateTokenPair(user User) (map[string]string, error) {
	// Create token
	// log.Println("Password is correct!")
	token := jwt.New(jwt.SigningMethodHS512)
	// Set claims
	// This is the information which frontend can use
	// The backend can also decode the token and get admin etc.
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["name"] = user.UserName
	claims["admin"] = user.IsAdmin
	// claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	for k, v := range claims {
		log.Printf("key = %v, type = %T, value = %v\n", k, v, v)
	}
	// Generate encoded token and send it as response.
	// The signing string should be secret (a generated UUID          works too)
	t, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS512)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = user.ID // sub is UID of the user
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// log.Printf("rtClaims   Type %T and expired at %[1]v", rtClaims["exp"])
	// log.Printf("time.Now() Type %T and expired at %[1]v", time.Now().Unix())

	rt, err := refreshToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  t,
		"refresh_token": rt,
	}, nil
}
