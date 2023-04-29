package user

import "time"

type User struct {
	ID        string
	Phone     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Response struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type RegisterUser struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UpdatePasswordUser struct {
	ID          string `json:"id"`
	Phone       string `json:"phone" validate:"required"`
	Password    string `json:"password" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}

type UpdatePhoneUser struct {
	ID       string `json:"id"`
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}
