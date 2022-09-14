package queries

import (
	"context"
	"github.com/ce-final-project/backend_rest_api/account_service/config"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/account/repository"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/models"
	"github.com/ce-final-project/backend_rest_api/pkg/logger"
)

type GetAccountByIdHandler interface {
	Handle(ctx context.Context, query *GetAccountByIdQuery) (*models.Account, error)
}

type getAccountByIdHandler struct {
	log    logger.Logger
	cfg    *config.Config
	pgRepo repository.Repository
}

func NewGetAccountByIdHandler(log logger.Logger, cfg *config.Config, pgRepo repository.Repository) *getAccountByIdHandler {
	return &getAccountByIdHandler{log: log, cfg: cfg, pgRepo: pgRepo}
}

func (q *getAccountByIdHandler) Handle(ctx context.Context, query *GetAccountByIdQuery) (*models.Account, error) {
	return q.pgRepo.GetAccountById(ctx, query.AccountID)
}
