package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: no matching record found")
var ErrInvalidCredentials = errors.New("models invalid credentials")
var ErrDuplicateEmail = errors.New("models: duplicate email")

type News struct {
	ID        int
	Title     string
	Content   string
	Image_url string
}

type Foods struct {
	ID        int
	Meal_name string
	Weekday   string
	Quantity  int
}

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Role           string
}
