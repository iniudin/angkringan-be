package entity

import "time"

type User struct {
	ID        int
	Phone     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
