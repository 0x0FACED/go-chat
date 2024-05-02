package utils

import (
	"errors"
	"go-chat/internal/models"
)

func ValidateUser(u *models.User) error {
	if len(u.Username) > 16 || len(u.Username) == 0 {
		return errors.New(ErrUsernameLength)
	}
	if len(u.Password) < 5 {
		return errors.New(ErrPasswordLength)
	}
	return nil
}
