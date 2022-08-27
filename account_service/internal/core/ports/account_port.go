package ports

import "github.com/ce-final-project/backend_rest_api/account_service/internal/core/domain"

type AccountRepository interface {
	GetAllAccount() ([]domain.Account, error)
	GetAccount(accountID uint64) (*domain.Account, error)
	StoreAccount(account *domain.Account) (*domain.Account, error)
	UpdateAccount(account *domain.Account) error
	DeleteAccount(accountID uint64) error
}

type AccountService interface {
	GetAllAccount() ([]domain.AccountResponse, error)
	GetAccount(accountID uint64) (*domain.AccountResponse, error)
	CreateAccount(account *domain.AccountRequest) (*domain.AccountResponse, error)
	ChangePassword(accountID uint64, oldPWD, newPWD string) error
	BanAccount(accountID uint64) (*domain.AccountResponse, error)
	DeleteAccount(accountID uint64) error
}
