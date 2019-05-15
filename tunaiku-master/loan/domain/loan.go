package domain

import (
	"time"

	"github.com/Charnojuntak/tunaiku/loan/repository"
)

func GetListOfTrackedLoan(date time.Time) (*repository.ListOfLoan, error) {
	where := "where date > $1 and date <= $2"
	startDate := date.AddDate(0, 0, -7)
	result, err := repository.GetListOfTrackedLoan(where, startDate, date)
	if err != nil {
		return result, err
	}

	return result, nil
}
