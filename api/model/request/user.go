package request

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
