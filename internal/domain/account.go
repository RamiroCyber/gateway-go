package domain

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Username  string
	Email     string
	APIKey    string
	balance   float64
	balanceMu sync.RWMutex
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GenerateAPIKey() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func NewAccount(username, email string) *Account {
	return &Account{
		ID:        uuid.New().String(),
		Username:  username,
		Email:     email,
		balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		APIKey:    GenerateAPIKey(),
	}
}

func (a *Account) AddBalance(amount float64) {
	a.balanceMu.Lock()
	defer a.balanceMu.Unlock()
	a.balance += amount
}

func (a *Account) GetBalance() float64 {
	a.balanceMu.RLock()
	defer a.balanceMu.RUnlock()
	return a.balance
}

func (a *Account) SetBalance(balance float64) {
	a.balanceMu.Lock()
	defer a.balanceMu.Unlock()
	a.balance = balance
}
