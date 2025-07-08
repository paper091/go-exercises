package main

import (
	"fmt"
)

func main() {
	// simulating an apis success case
	resp := responseStruct{
		status:  200,
		message: "SUCCESS",
		data: userData{
			id:    "user123",
			name:  "user 123",
			email: "user123@gmail.com",
			role:  "admin",
		},
		statusCode: 200,
	}

	fmt.Println(resp)
}

type responseStruct struct {
	status     int
	statusCode int
	message    string
	data       userData
}

type userData struct {
	id    string
	name  string
	email string
	role  string
}
