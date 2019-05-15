package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Charnojuntak/tunaiku/database"
)

type ListOfLoan struct {
	Count   int     `db:"count" json:"count"`
	Summary int     `db:"summary" json:"summary"`
	Average float64 `db:"average" json:"average"`
}

const (
	queryListOfLoan = `select count(user_id) as count, coalesce(sum(amount), 0) as summary, coalesce(avg(amount), 0) as average from users `
)

func GetListOfTrackedLoan(where string, args ...interface{}) (*ListOfLoan, error) {
	var result ListOfLoan
	query := fmt.Sprintf("%s%s", queryListOfLoan, where)
	fmt.Println(query)
	err := database.Client.Get(&result, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return &result, errors.New("there no data")
		}
		return &result, err
	}

	return &result, nil
}
