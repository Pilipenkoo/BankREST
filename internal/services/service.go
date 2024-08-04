package services

import (
	"BankRESTAPI/internal/account"
	"errors"
	"sync"
)

// Service provides account operations
type Service struct {
	accounts   map[string]*account.Account
	mu         sync.Mutex
	createCh   chan *account.Account
	depositCh  chan depositRequest
	withdrawCh chan withdrawRequest
	balanceCh  chan balanceRequest
}

// NewService creates a new Service
func NewService() *Service {
	s := &Service{
		accounts:   make(map[string]*account.Account),
		createCh:   make(chan *account.Account),
		depositCh:  make(chan depositRequest),
		withdrawCh: make(chan withdrawRequest),
		balanceCh:  make(chan balanceRequest),
	}
	go s.run()
	return s
}

type depositRequest struct {
	id     string
	amount float64
	result chan error
}

type withdrawRequest struct {
	id     string
	amount float64
	result chan error
}

type balanceRequest struct {
	id     string
	result chan balanceResponse
}

type balanceResponse struct {
	balance float64
	err     error
}

func (s *Service) run() {
	for {
		select {
		case acc := <-s.createCh:
			s.mu.Lock()
			s.accounts[acc.ID] = acc
			s.mu.Unlock()
		case req := <-s.depositCh:
			s.mu.Lock()
			acc, exists := s.accounts[req.id]
			s.mu.Unlock()
			if !exists {
				req.result <- errors.New("account not found")
			} else {
				req.result <- acc.Deposit(req.amount)
			}
		case req := <-s.withdrawCh:
			s.mu.Lock()
			acc, exists := s.accounts[req.id]
			s.mu.Unlock()
			if !exists {
				req.result <- errors.New("account not found")
			} else {
				req.result <- acc.Withdraw(req.amount)
			}
		case req := <-s.balanceCh:
			s.mu.Lock()
			acc, exists := s.accounts[req.id]
			s.mu.Unlock()
			if !exists {
				req.result <- balanceResponse{0, errors.New("account not found")}
			} else {
				req.result <- balanceResponse{acc.GetBalance(), nil}
			}
		}
	}
}

func (s *Service) CreateAccount() *account.Account {
	acc := account.NewAccount(generateID())
	s.createCh <- acc
	return acc
}

func (s *Service) Deposit(id string, amount float64) error {
	result := make(chan error)
	s.depositCh <- depositRequest{id, amount, result}
	return <-result
}

func (s *Service) Withdraw(id string, amount float64) error {
	result := make(chan error)
	s.withdrawCh <- withdrawRequest{id, amount, result}
	return <-result
}

func (s *Service) GetBalance(id string) (float64, error) {
	result := make(chan balanceResponse)
	s.balanceCh <- balanceRequest{id, result}
	res := <-result
	return res.balance, res.err
}
