package main

import (
	"fmt"
	"log"

	bcrypt "golang.org/x/crypto/bcrypt"
)

func main() {
	userPassword1 := "Welcome1"

	// Generate "hash" to store from user password
	hash, err := bcrypt.GenerateFromPassword([]byte(userPassword1), bcrypt.DefaultCost+1)
	if err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
	}
	fmt.Printf("Hash to store: %q - len %d\n", string(hash), len(hash))
	// Store this "hash" somewhere, e.g. in your database

	// After a while, the user wants to log in and you need to check the password he entered
	userPassword2 := "Welcome1"
	hashFromDatabase := string(hash)

	// Comparing the password with the hash
	if err := bcrypt.CompareHashAndPassword([]byte(hashFromDatabase), []byte(userPassword2)); err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
	}

	fmt.Println("Password was correct!")
}
