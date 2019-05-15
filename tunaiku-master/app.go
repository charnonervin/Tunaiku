package main

import (
	"log"
	"net/http"

	"github.com/Charnojuntak/tunaiku/database"
	"github.com/julienschmidt/httprouter"

	cicilanHandler "github.com/Charnojuntak/tunaiku/cicilan/handler"
	loanHandler "github.com/Charnojuntak/tunaiku/loan/handler"
	userHandler "github.com/Charnojuntak/tunaiku/user/handler"
)

func main() {
	database.InitDB()

	router := httprouter.New()
	userHandler.RegisterHandler(router)
	loanHandler.RegisterHandler(router)
	cicilanHandler.RegisterHandler(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
