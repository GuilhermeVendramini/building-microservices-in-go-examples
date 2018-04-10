package main

import (
	"errors"
)

// UserService provides users operations.
type UserService interface {
	Get(string) (string, error)
	Delete(int) (bool, error)
}

type userService struct{}

func (userService) Get(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	/*
		Your code here
	*/
	return "Guilherme", nil
}

func (userService) Delete(s int) (bool, error) {
	/*
		Your code here
	*/
	return true, nil
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")
