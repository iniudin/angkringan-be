package entity

import "time"

type User struct {
	ID        string
	Phone     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
