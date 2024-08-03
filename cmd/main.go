package main

import (
	"BankRESTAPI/internal/api"
	"BankRESTAPI/internal/services"
	"log"
	"net/http"
)

func main() {
	accountService := services.NewService()
	handlerAccount := api.NewHandler(accountService)
	router := api.NewRouter(handlerAccount)

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
