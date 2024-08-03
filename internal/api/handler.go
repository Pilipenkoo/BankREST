package api

import (
	"BankRESTAPI/internal/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct {
	accountService *services.Service
}

func NewHandler(accountService *services.Service) *Handler {
	return &Handler{
		accountService: accountService,
	}
}

func (h *Handler) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	acc := h.accountService.CreateAccount()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(acc)
}

func (h *Handler) DepositHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	amountStr := r.URL.Query().Get("amount")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.accountService.Deposit(id, amount)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	amountStr := r.URL.Query().Get("amount")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	err = h.accountService.Withdraw(id, amount)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
}

// GetBalanceHandler handles balance inquiries
func (h *Handler) GetBalanceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	balance, err := h.accountService.GetBalance(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(balance)
}
