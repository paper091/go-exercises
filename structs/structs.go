package main

import (
	"fmt"
	"time"
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
		timestamp:  time.Now(),
	}

	user := resp.data
	// inline structure
	tempData := struct {
		eldapId        string
		eldapUserName  string
		systemUserInfo userData
	}{
		"1221000",
		"user123eldap",
		user,
	}

	fmt.Println(resp)
	fmt.Println(tempData)
}

type responseStruct struct {
	status     int
	statusCode int
	message    string
	data       userData // struct embedding
	timestamp  time.Time
}

type userData struct {
	id    string
	name  string
	email string
	role  string
}
