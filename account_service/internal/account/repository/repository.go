package repository

import (
	"context"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/models"
	"github.com/ce-final-project/backend_rest_api/pkg/utils"
	uuid "github.com/satori/go.uuid"
)

type Repository interface {
	CreateAccount(ctx context.Context, account *models.Account) (*models.Account, error)
	UpdateAccount(ctx context.Context, account *models.Account) (*models.Account, error)
	DeleteAccount(ctx context.Context, uuid uuid.UUID) error

	GetAccountById(ctx context.Context, uuid uuid.UUID) (*models.Account, error)
	Search(ctx context.Context, search string, pagination *utils.Pagination) (*models.AccountsList, error)
}

type CacheRepository interface {
	PutAccount(ctx context.Context, key string, account *models.Account)
	GetAccount(ctx context.Context, key string) (*models.Account, error)
	DelAccount(ctx context.Context, key string)
	DelAllAccounts(ctx context.Context)
}
