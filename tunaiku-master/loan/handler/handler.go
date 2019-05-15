package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Charnojuntak/tunaiku/loan/domain"
	"github.com/julienschmidt/httprouter"
)

func RegisterHandler(router *httprouter.Router) {
	router.GET("/loan/listoftrackedloan", GetListOfTrackedLoan)
}

func GetListOfTrackedLoan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	date, _ := time.Parse("2006-01-02", r.FormValue("date"))

	listOfTrackedLoan, err := domain.GetListOfTrackedLoan(date)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	result, err := json.Marshal(listOfTrackedLoan)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(result)
	return
}
