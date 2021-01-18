package model

type LoginParam struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterParam struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password" validate:"required"`
}
