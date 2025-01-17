package storage

import "authService/internal/models"

type Storage interface {
	GetUserById() (models.User, error)
	GetUserByUsername() (models.User, error)
	GetUserByEmail() (models.User, error)
	CreateUser(models.User) error
	UpdateUser(models.User) error
	DeleteUser(models.User) error
}
