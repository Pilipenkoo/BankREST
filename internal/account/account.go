package account

import (
	"errors"
	"log"
	"sync"
)

type Account struct {
	ID      string
	balance float64
	mu      sync.Mutex
}

func NewAccount(id string) *Account {
	return &Account{
		ID: id,
	}
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	} else {
		a.mu.Lock()
		defer a.mu.Unlock()
		a.balance += amount
		log.Printf("Deposit: AccountID=%s, Amount=%f, NewBalance=%f", a.ID, amount, a.balance)
		return nil
	}
}

func (a *Account) Withdraw(amount float64) error {
	if amount > a.balance {
		return errors.New("amount must be less than balance")
	} else {
		a.mu.Lock()
		defer a.mu.Unlock()
		a.balance -= amount
		log.Printf("Deposit: AccountID=%s, Amount=%f, NewBalance=%f", a.ID, amount, a.balance)
		return nil
	}
}

func (a *Account) GetBalance() float64 {
	a.mu.Lock()
	defer a.mu.Unlock()

	return a.balance
}
