package repository

import (
	"context"
	"github.com/ce-final-project/backend_rest_api/account_service/config"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/models"
	"github.com/ce-final-project/backend_rest_api/pkg/logger"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type accountRepository struct {
	log logger.Logger
	cfg *config.Config
	db  *pgxpool.Pool
}

func NewAccountRepository(log logger.Logger, cfg *config.Config, db *pgxpool.Pool) *accountRepository {
	return &accountRepository{log: log, cfg: cfg, db: db}
}

func (p *accountRepository) CreateAccount(ctx context.Context, account *models.Account) (*models.Account, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "accountRepository.CreateAccount")
	defer span.Finish()

	var created models.Account
	if err := p.db.QueryRow(ctx, createAccountQuery, &account.AccountID, &account.PlayerID, &account.Username, &account.Email, &account.PasswordHash, &account.IsBan).Scan(
		&created.AccountID,
		&created.PlayerID,
		&created.Username,
		&created.Email,
		&created.PasswordHash,
		&created.IsBan,
		&created.CreatedAt,
		&created.UpdatedAt,
	); err != nil {
		return nil, errors.Wrap(err, "db.QueryRow")
	}

	return &created, nil
}

func (p *accountRepository) UpdateAccount(ctx context.Context, account *models.Account) (*models.Account, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "accountRepository.UpdateAccount")
	defer span.Finish()

	var prod models.Account
	if err := p.db.QueryRow(
		ctx,
		updateAccountQuery,
		&account.Username,
		&account.Email,
		&account.PasswordHash,
		&account.IsBan,
		&account.AccountID,
	).Scan(&prod.AccountID, &prod.PlayerID, &prod.Username, &prod.Email, &prod.PasswordHash, &prod.IsBan, &prod.CreatedAt, &prod.UpdatedAt); err != nil {
		return nil, errors.Wrap(err, "Scan")
	}

	return &prod, nil
}

func (p *accountRepository) GetAccountById(ctx context.Context, uuid uuid.UUID) (*models.Account, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "accountRepository.GetAccountById")
	defer span.Finish()

	var account models.Account
	if err := p.db.QueryRow(ctx, getAccountByIdQuery, uuid).Scan(
		&account.AccountID,
		&account.PlayerID,
		&account.Username,
		&account.Email,
		&account.PasswordHash,
		&account.IsBan,
		&account.CreatedAt,
		&account.UpdatedAt,
	); err != nil {
		return nil, errors.Wrap(err, "Scan")
	}

	return &account, nil
}

func (p *accountRepository) DeleteAccountByID(ctx context.Context, uuid uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "accountRepository.DeleteAccountByID")
	defer span.Finish()

	_, err := p.db.Exec(ctx, deleteAccountByIdQuery, uuid)
	if err != nil {
		return errors.Wrap(err, "Exec")
	}

	return nil
}
