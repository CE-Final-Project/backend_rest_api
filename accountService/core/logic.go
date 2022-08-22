package core

import (
	"errors"
	"fmt"
	"github.com/teris-io/shortid"
	"gopkg.in/dealancer/validate.v2"
	"time"
)

var (
	ErrAccountNotFound = errors.New("account not found")
	ErrAccountInvalid  = errors.New("account invalid")
)

type accountService struct {
	accountRepo AccountRepository
}

func NewAccountService(accountRepo AccountRepository) AccountService {
	return &accountService{
		accountRepo,
	}
}

func (a *accountService) Find(playerId string, accountId string) (*Account, error) {
	return a.accountRepo.Find(playerId, accountId)
}

func (a *accountService) Store(account *Account) error {
	if err := validate.Validate(account); err != nil {
		return fmt.Errorf("%v service.account.store", ErrAccountInvalid)
	}
	account.PlayerId = shortid.MustGenerate()
	account.CreatedAt = time.Now().UTC().Unix()
	return a.accountRepo.Store(account)
}

func (a *accountService) Remove(playerId string, accountId string) (*Account, error) {
	return a.accountRepo.Remove(playerId, accountId)
}