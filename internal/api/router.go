package api

import (
	"github.com/gorilla/mux"
)

func NewRouter(handler *Handler) *mux.Router {
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/accounts", handler.CreateAccountHandler).Methods("POST")
	router.HandleFunc("/accounts/{id}/deposit", handler.DepositHandler).Methods("POST")
	router.HandleFunc("/accounts/{id}/withdraw", handler.WithdrawHandler).Methods("POST")
	router.HandleFunc("/accounts/{id}/balance", handler.GetBalanceHandler).Methods("GET")
	return router
}
