package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID               int       `json:"id" db:"id"`
	Name             string    `json:"name" db:"name"`
	Username         string    `json:"username" db:"username"`
	Password         string    `json:"password" db:"password"`
	RegistrationDate time.Time `json:"reg_date" db:"registration_date"`
	Description      string    `json:"desc" db:"description"`
}

func (u *User) HashPassword(pass string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return pass, err
	}
	return string(hashed), nil
}

func (u *User) CompareHashPassword(pass string, hashed string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pass))
	if err != nil {
		return err
	}
	return nil
}
