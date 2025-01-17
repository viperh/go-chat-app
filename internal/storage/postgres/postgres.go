package postgres

import (
	"authService/internal/config"
	"authService/internal/models"
	"authService/internal/rlog"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Database *gorm.DB
	Logger   *rlog.Logger
}

func NewDatabase(cfg *config.Config) *Postgres {
	logger := rlog.NewLogger(cfg.LogLevel, "POSTGRES")
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName, cfg.DbSslMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("Could not connect to postgres!")
	}

	return &Postgres{
		Database: db,
		Logger:   logger,
	}
}

func (p *Postgres) GetUserById(id int) (*models.User, error) {

	var user models.User

	result := p.Database.First(&user, id)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &user, nil
}

func (p *Postgres) GetUserByUsername(username string) (*models.User, error) {
	var user models.User

	result := p.Database.First(&user, "username = ?", username)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &user, nil
}

func (p *Postgres) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	result := p.Database.First(&user, "email = ?", email)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &user, nil
}

func (p *Postgres) CreateUser(user *models.User) error {
	result := p.Database.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *Postgres) UpdateUser(user *models.User) error {
	result := p.Database.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *Postgres) DeleteUser(id int) error {
	var user models.User
	result := p.Database.Delete(&user, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
