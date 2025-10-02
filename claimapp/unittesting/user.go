package main

import (
	"errors"
	"regexp"
)

type User struct {
	Name     string
	Email    string
	Password string
}

var (
	ErrInvalidUserName  = errors.New("invalid user name")
	ErrInvalidUserEmail = errors.New("invalid user email")
	ErrInvalidUserPass  = errors.New("invalid user password")
)

func ValidateUser(u User) error {

	// user name not empty
	if u.Name == "" {
		return ErrInvalidUserName
	}
	if u.Email == "" {
		return ErrInvalidUserEmail
	}
	if u.Password == "" {
		return ErrInvalidUserPass
	}

	//user name regular expression
	re := regexp.MustCompile(`^[a-zA-Z]{3,20}$`)
	if !re.MatchString(u.Name) {
		return ErrInvalidUserName
	}
	//user email regular expression
	re = regexp.MustCompile(`^[a-zA-Z0-9._%+-]{1,}@[a-zA-Z0-9.-]{2,}\.[a-zA-Z]{2,}$`)
	if !re.MatchString(u.Email) {
		return ErrInvalidUserEmail
	}
	// user password validation: At least 8 characters, one uppercase, one lowercase, one number, one special character
	if len(u.Password) < 8 {
		return ErrInvalidUserPass
	}
	reLower := regexp.MustCompile(`[a-z]`)
	reUpper := regexp.MustCompile(`[A-Z]`)
	reDigit := regexp.MustCompile(`\d`)
	reSpecial := regexp.MustCompile(`[@$!%*?&]`)
	if !reLower.MatchString(u.Password) ||
		!reUpper.MatchString(u.Password) ||
		!reDigit.MatchString(u.Password) ||
		!reSpecial.MatchString(u.Password) {
		return ErrInvalidUserPass
	}
	return nil
}
