package services

import (
	"github.com/ce-final-project/backend_rest_api/account_service/internal/core/domain"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/core/ports"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/pkg/password"
	"github.com/teris-io/shortid"
	"log"
	"time"
)

type accountService struct {
	accRepo ports.AccountRepository
}

func (a *accountService) GetAllAccount() ([]domain.AccountResponse, error) {
	var accountsRes []domain.AccountResponse
	accounts, err := a.accRepo.GetAllAccount()
	if err != nil {
		log.Printf("Retriving accounts error %v", err)
		return nil, err
	}
	for _, account := range accounts {
		accountRes := domain.AccountResponse{
			AccountID: account.AccountID,
			PlayerID:  account.PlayerID,
			Username:  account.Username,
			Email:     account.Email,
			IsBan:     account.IsBan,
			CreatedAt: account.CreatedAt,
		}
		accountsRes = append(accountsRes, accountRes)
	}
	return accountsRes, nil
}

func (a *accountService) GetAccount(accountID uint64) (*domain.AccountResponse, error) {
	account, err := a.accRepo.GetAccount(accountID)
	if err != nil {
		log.Printf("Retriving account error %v", err)
		return nil, err
	}
	return &domain.AccountResponse{
		AccountID: account.AccountID,
		PlayerID:  account.PlayerID,
		Username:  account.Username,
		Email:     account.Email,
		IsBan:     account.IsBan,
		CreatedAt: account.CreatedAt,
	}, nil
}

func (a *accountService) CreateAccount(account *domain.AccountRequest) (*domain.AccountResponse, error) {
	playerId := shortid.MustGenerate()
	passwordHash, err := password.HashPassword(account.Password)
	if err != nil {
		log.Printf("Hashing password error %v", err)
		return nil, err
	}
	newAccount := &domain.Account{
		PlayerID:     playerId,
		Username:     account.Username,
		Email:        account.Email,
		PasswordHash: passwordHash,
		IsBan:        false,
		CreatedAt:    time.Now().Unix(),
	}
	var accCreated *domain.Account
	accCreated, err = a.accRepo.StoreAccount(newAccount)
	if err != nil {
		log.Printf("Create Account error %v", err)
		return nil, err
	}
	return &domain.AccountResponse{
		AccountID: accCreated.AccountID,
		PlayerID:  accCreated.PlayerID,
		Username:  accCreated.Username,
		Email:     accCreated.Email,
		IsBan:     accCreated.IsBan,
		CreatedAt: accCreated.CreatedAt,
	}, nil
}

func (a *accountService) ChangePassword(accountID uint64, oldPWD, newPWD string) error {
	account, err := a.accRepo.GetAccount(accountID)
	if err != nil {
		log.Printf("Retriving account error %v", err)
		return err
	}
	pwdCorrect := password.CheckPasswordHash(oldPWD, account.PasswordHash)
	if !pwdCorrect {
		log.Printf("Old Password is not correct")
		return nil
	}
	var pwdHashed string
	pwdHashed, err = password.HashPassword(newPWD)
	if err != nil {
		return err
	}
	account.PasswordHash = pwdHashed
	err = a.accRepo.UpdateAccount(account)
	return err
}

func (a *accountService) BanAccount(accountID uint64) (*domain.AccountResponse, error) {
	account, err := a.accRepo.GetAccount(accountID)
	if err != nil {
		log.Printf("Retriving account error %v", err)
		return nil, err
	}
	account.IsBan = true
	err = a.accRepo.UpdateAccount(account)
	return &domain.AccountResponse{
		AccountID: account.AccountID,
		PlayerID:  account.PlayerID,
		Username:  account.Username,
		Email:     account.Email,
		IsBan:     account.IsBan,
		CreatedAt: account.CreatedAt,
	}, nil
}

func (a *accountService) DeleteAccount(accountID uint64) error {
	err := a.accRepo.DeleteAccount(accountID)
	if err != nil {
		return err
	}
	return nil
}

func NewAccountService(accRepo ports.AccountRepository) ports.AccountService {
	return &accountService{accRepo}
}
