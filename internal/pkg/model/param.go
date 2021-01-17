package model

type LoginParam struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RegisterParam struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
