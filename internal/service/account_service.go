package service

import (
	"errors"
	"fmt"

	"github.com/RamiroCyber/gateway-go/internal/constants"
	"github.com/RamiroCyber/gateway-go/internal/domain"
	"github.com/RamiroCyber/gateway-go/internal/dto"
	"github.com/RamiroCyber/gateway-go/internal/repository"
)

var (
	ErrFailedToGenerateUniqueAPIKey = errors.New("failed to generate unique API key after maximum attempts")
)

type AccountService struct {
	repository *repository.AccountRepository
}

func NewAccountService(repository *repository.AccountRepository) *AccountService {
	return &AccountService{repository: repository}
}

func (s *AccountService) CreateAccount(input dto.AccountInput) (*dto.AccountOutput, error) {
	account := dto.ToAccount(input)

	for i := 0; i < constants.MaxAPIKeyGenerationAttempts; i++ {
		existingAccount, err := s.repository.FindByAPIKey(account.APIKey)
		if err != nil && err != domain.ErrAccountNotFound {
			return nil, fmt.Errorf("error checking API key existence: %w", err)
		}

		if existingAccount == nil {
			if err := s.repository.Save(account); err != nil {
				return nil, fmt.Errorf("error saving account: %w", err)
			}

			output := dto.FromAccount(account)
			return &output, nil
		}

		account.APIKey = domain.GenerateAPIKey()
	}

	return nil, ErrFailedToGenerateUniqueAPIKey
}

func (s *AccountService) GetAccountByAPIKey(apiKey string) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}
	output := dto.FromAccount(account)
	return &output, nil
}

func (s *AccountService) GetAccountByID(id string) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	output := dto.FromAccount(account)
	return &output, nil
}

func (s *AccountService) UpdateBalance(apiKey string, amount float64) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	account.AddBalance(amount)
	err = s.repository.UpdateBalance(account)
	if err != nil {
		return nil, err
	}
	output := dto.FromAccount(account)
	return &output, nil
}
