package util

import (
	"regexp"
	"errors"
)

/** Валидация логина пользователя. */
func ValidateLogin(login string) error {
	if regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString(login) {
		return nil
	} else {
		return errors.New("Incorrect login")
	}
}

/** Валидация логина пользователя. */
func ValidatePassword(password string) error {
	if password != "" {
		return nil
	} else {
		return errors.New("Incorrect password")
	}
}