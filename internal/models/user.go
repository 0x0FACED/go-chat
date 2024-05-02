package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID               int       `json:"id" db:"id"`
	Name             string    `json:"name" db:"name"`
	Username         string    `json:"username" db:"username"`
	Password         string    `json:"password" db:"password"`
	RegistrationDate time.Time `json:"registration_date" db:"registration_date"`
	Description      string    `json:"description" db:"description"`
}

func (u *User) HashPassword(pass string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	u.Password = string(hashed)
	return err
}
