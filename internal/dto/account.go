package dto

import (
	"time"

	"github.com/RamiroCyber/gateway-go/internal/domain"
)

type AccountInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type AccountOutput struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Balance   float64   `json:"balance"`
	APIKey    string    `json:"api_key,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToAccount(input AccountInput) *domain.Account {
	return domain.NewAccount(input.Username, input.Email)
}

func FromAccount(account *domain.Account) AccountOutput {
	return AccountOutput{
		ID:        account.ID,
		Username:  account.Username,
		Email:     account.Email,
		Balance:   account.GetBalance(),
		APIKey:    account.APIKey,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
}
