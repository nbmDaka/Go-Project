package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: no matching record found")

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
