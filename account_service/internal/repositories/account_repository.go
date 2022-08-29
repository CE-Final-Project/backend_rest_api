package repositories

import (
	"errors"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/core/domain"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/core/ports"
	"github.com/jmoiron/sqlx"
)

type accountRepository struct {
	db *sqlx.DB
}

func (a *accountRepository) StoreAccount(account *domain.Account) (*domain.Account, error) {
	query := "INSERT INTO accounts ( player_id, username, email, password_hash, is_ban, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING account_id"
	err := a.db.QueryRow(query, account.PlayerID, account.Username, account.Email, account.PasswordHash, account.IsBan, account.CreatedAt).Scan(&account.AccountID)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (a *accountRepository) GetAllAccount() ([]domain.Account, error) {
	limit := 10
	offset := 0
	var accounts []domain.Account
	query := "SELECT account_id, player_id, username, email, password_hash, is_ban, created_at FROM accounts LIMIT $1 OFFSET $2"
	err := a.db.Select(&accounts, query, limit, offset)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (a *accountRepository) GetAccount(accountID uint64) (*domain.Account, error) {
	var account domain.Account
	query := "SELECT account_id, player_id, username, email, password_hash, is_ban, created_at FROM accounts WHERE account_id=$1 LIMIT 1"
	err := a.db.Get(&account, query, accountID)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (a *accountRepository) UpdateAccount(account *domain.Account) error {
	query := "UPDATE accounts SET username=$1, password_hash=$2, is_ban=$3 WHERE account_id=$4"
	_, err := a.db.Exec(query, account.Username, account.PasswordHash, account.IsBan, account.AccountID)
	if err != nil {
		return err
	}
	return nil
}

func (a *accountRepository) DeleteAccount(accountID uint64) error {
	query := "DELETE FROM accounts WHERE account_id=$1"
	result, err := a.db.Exec(query, accountID)
	if err != nil {
		return err
	}
	count, rowErr := result.RowsAffected()
	if rowErr != nil {
		return rowErr
	}
	if count == 0 {
		return errors.New("account not found account.repository.DeleteAccount")
	}
	return nil
}

func NewAccountRepository(db *sqlx.DB) ports.AccountRepository {
	return &accountRepository{db}
}
