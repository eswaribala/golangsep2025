package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FetchUser(apiURL string) (User, error) {
	resp, err := http.Get(apiURL + "/users/1")
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	var user User
	err = json.NewDecoder(resp.Body).Decode(&user)
	return user, err
}
