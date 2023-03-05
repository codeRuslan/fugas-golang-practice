package main

import (
	"fmt"
	"log"
	"strings"
)

type user struct {
	name  string
	age   uint8
	email string
}

var users []user

func main() {
	users = []user{
		{
			name:  "Mike",
			age:   32,
			email: "mike@gmail.com",
		},
		{
			name:  "John",
			age:   54,
			email: "john@gmail.com",
		},
		{
			name:  "Abobus",
			age:   2,
			email: "abobus@gmail.com",
		},
	}

	userName, userAge, userEmail := userFinder()
	fmt.Printf("Name: %s, Age: %d, Email: %s\n", userName, userAge, userEmail)
}

func userFinder() (name string, age uint8, email string) {
	var inputName string
	fmt.Scan(&inputName)
	inputName = strings.Title(inputName[:1]) + strings.ToLower(inputName[1:])

	for n := 0; n < len(users); n++ {
		if users[n].name == inputName {
			name = users[n].name
			age = users[n].age
			email = users[n].email
			return name, age, email
		}
	}

	log.Println("Warning: No user have been found. Returning standard empty user")
	return "NIL", 0, "NIL"
}
