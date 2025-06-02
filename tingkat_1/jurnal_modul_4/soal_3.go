package main

import (
	"fmt"
)

func loginValid(username, password string) bool {
	const usernameBenar = "telu"
	const passwordBenar = "telu1234"

	return username == usernameBenar && password == passwordBenar
}

func main3() {
	kasusUji := [][]string{
		{"TELU", "2024"},        
		{"telu", "telu"},        
		{"telu", "telu1234"},   
		{"Telu", "telu1234"},   
	}

	for _, kasus := range kasusUji {
		username := kasus[0]
		password := kasus[1]
		hasil := loginValid(username, password)
		fmt.Printf("Username: %s, Password: %s, Login Valid: %v\n", username, password, hasil)
	}
}