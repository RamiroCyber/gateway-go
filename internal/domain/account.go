package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Username  string
	Email     string
	APIKey    string
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(username, email string) *Account {
	return &Account{
		ID:        uuid.New().String(),
		Username:  username,
		Email:     email,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (a *Account) Validate() error {
	if a.Username == "" {
		return errors.New("username is required")
	}
	if a.Email == "" {
		return errors.New("email is required")
	}
	if a.APIKey == "" {
		return errors.New("api key is required")
	}
	if a.Balance < 0 {
		return errors.New("balance cannot be negative")
	}
	return nil
}

func (a *Account) GenerateAPIKey() string {
	return uuid.New().String()
}
