package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	password := "mypassword123"

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Hashed password:", string(hash))
}