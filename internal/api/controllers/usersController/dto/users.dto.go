package dto

import "authService/internal/models"

type User struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type ID struct {
	ID int `json:"id"`
}

type Username struct {
	Username string `json:"username"`
}

type Email struct {
	Email string `json:"email"`
}

type ReturnUser struct {
	User  *models.User
	Error error
}

type ReturnConfirmation struct {
	Confirmation bool
	Error        error
}
