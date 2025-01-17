package models

type User struct {
	ID        int    `gorm:"primary_key"`
	Firstname string `gorm:"not null"`
	Lastname  string `gorm:"not null"`
	Username  string `gorm:"unique"`
	Email     string `gorm:"unique"`
	Password  string `gorm:"not null"`
}
