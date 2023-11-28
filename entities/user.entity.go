package entities

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Nip       string    `json:"nip" validate:"required"`
	Firstname string    `json:"firstname" validate:"required"`
	Lastname  string    `json:"lastname" validate:"required"`
	Email     string    `json:"email" validate:"email"`
	Image     string    `json:"image"`
	Password  string    `json:"password" validate:"required"`
	Token     string    `json:"token"`
	LastLogin time.Time `json:"last_login"`
	IsActive  int8      `json:"is_active"`
}

func (e *User) TableName() string {
	return "user"
}

func (e *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	e.Password = string(bytes)
	return nil
}
