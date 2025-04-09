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
	balanceMu sync.Mutex
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(username, email string) *Account {
	return &Account{
		ID:        uuid.New().String(),
		Username:  username,
		Email:     email,
		balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		APIKey:    generateAPIKey(),
	}
}

func generateAPIKey() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func (a *Account) AddBalance(amount float64) {
	a.balanceMu.Lock()
	defer a.balanceMu.Unlock()
	a.balance += amount
}

func (a *Account) SubtractBalance(amount float64) {
	a.balanceMu.Lock()
	defer a.balanceMu.Unlock()
	a.balance -= amount
}
