package repository

import (
	"time"

	"github.com/Charnojuntak/tunaiku/database"
)

type User struct {
	Date           time.Time `db:"date"`
	DateParam      string    `json:"date"`
	KTP            string    `db:"ktp" json:"ktp"`
	BirthDate      time.Time `db:"birth_date"`
	BirthDateParam string    `json:"birth_date"`
	Gender         string    `db:"gender" json:"gender"`
	Name           string    `db:"name" json:"name"`
	Amount         int       `db:"amount" json:"amount"`
	Period         int       `db:"period" json:"period"`
}

const (
	userInsertQuery = `insert into users (date, ktp, birth_date, gender, name, amount, period) values (:date, :ktp, :birth_date, :gender, :name, :amount, :period)`
)

func InsertUserData(user *User) error {
	_, err := database.Client.NamedExec(userInsertQuery, user)
	if err != nil {
		return err
	}

	return err
}
