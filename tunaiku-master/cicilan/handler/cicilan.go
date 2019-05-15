package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Charnojuntak/tunaiku/cicilan/domain"
	"github.com/julienschmidt/httprouter"
)

func RegisterHandler(router *httprouter.Router) {
	router.POST("/cicilan/simulate", SimulateCicilan)
}

func SimulateCicilan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var param domain.CicilanParam
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = json.Unmarshal([]byte(body), &param)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	param.Date, _ = time.Parse("2006-01-02", param.DateParam)

	cicilanResponse := domain.SimulateCicilan(&param)
	result, _ := json.Marshal(cicilanResponse)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(result)
	return
}
