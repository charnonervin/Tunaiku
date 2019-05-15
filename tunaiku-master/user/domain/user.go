package domain

import (
	"errors"

	"github.com/Charnojuntak/tunaiku/user/repository"
)

const (
	ErrKTPNumberNotValid = "ktp number isn't valid"
)

func AddUserData(user *repository.User) error {
	isValid := validateKtp(user.KTP)
	if !isValid {
		return errors.New(ErrKTPNumberNotValid)
	}

	err := repository.InsertUserData(user)
	if err != nil {
		return err
	}

	return nil
}

func validateKtp(ktpNumber string) bool {
	return true
}
