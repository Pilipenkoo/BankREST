package services

import (
	"BankRESTAPI/internal/account"
	"errors"
	"sync"
)

type Service struct {
	accounts map[string]*account.Account
	mu       sync.Mutex
}

func NewService() *Service {
	return &Service{
		accounts: make(map[string]*account.Account),
	}
}

func (s *Service) CreateAccount() *account.Account {
	s.mu.Lock()
	defer s.mu.Unlock()
	acc := account.NewAccount(generateID())
	s.accounts[acc.ID] = acc
	return acc
}

func (s *Service) Deposit(id string, amount float64) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	acc, exists := s.accounts[id]
	if !exists {
		return errors.New("account does not exist")
	} else {
		acc.Deposit(amount)
		return nil
	}
}

func (s *Service) Withdraw(id string, amount float64) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	acc, exists := s.accounts[id]
	if !exists {
		return errors.New("account does not exist")
	} else {
		acc.Withdraw(amount)
		return nil
	}
}

func (s *Service) GetBalance(id string) (float64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	acc, exists := s.accounts[id]
	if !exists {
		return 0, errors.New("account does not exist")
	} else {
		return acc.GetBalance(), nil
	}
}
