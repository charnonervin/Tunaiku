package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Charnojuntak/tunaiku/user/domain"
	"github.com/Charnojuntak/tunaiku/user/repository"
	"github.com/julienschmidt/httprouter"
)

type UserResponse struct {
	KTP    string `json:"ktp"`
	Status string `json:"status"`
}

func RegisterHandler(router *httprouter.Router) {
	router.POST("/user/addUser", AddUser)
}

func AddUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	response := UserResponse{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	var param repository.User
	err = json.Unmarshal([]byte(body), &param)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	param.Date, _ = time.Parse("2006-01-02", param.DateParam)
	param.BirthDate, _ = time.Parse("2006-01-02", param.BirthDateParam)
	err = domain.AddUserData(&param)
	if err != nil {
		if err.Error() == domain.ErrKTPNumberNotValid {
			response.KTP = param.KTP
			response.Status = "Invalid"
			result, _ := json.Marshal(response)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(result)
		} else {
			w.Write([]byte(err.Error()))
			return
		}
	}

	response.KTP = param.KTP
	response.Status = "Valid"
	result, _ := json.Marshal(response)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(result)
	return
}
