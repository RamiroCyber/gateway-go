package domain

type AccountRepository interface {
	FindByAPIKey(apiKey string) (*Account, error)
	FindByID(id string) (*Account, error)
	UpdateBalance(*Account) error
	Update(account *Account) error
	Save(account *Account) error
}
