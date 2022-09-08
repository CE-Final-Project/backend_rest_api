package queries

import (
	"context"
	"github.com/ce-final-project/backend_rest_api/account_service/config"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/account/repository"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/models"
	"github.com/ce-final-project/backend_rest_api/pkg/logger"
	"github.com/opentracing/opentracing-go"
)

type GetAccountByIdHandler interface {
	Handle(ctx context.Context, query *GetAccountByIdQuery) (*models.Account, error)
}

type getAccountByIdHandler struct {
	log       logger.Logger
	cfg       *config.Config
	mongoRepo repository.Repository
	redisRepo repository.CacheRepository
}

func NewGetAccountByIdHandler(log logger.Logger, cfg *config.Config, mongoRepo repository.Repository, redisRepo repository.CacheRepository) *getAccountByIdHandler {
	return &getAccountByIdHandler{log: log, cfg: cfg, mongoRepo: mongoRepo, redisRepo: redisRepo}
}

func (q *getAccountByIdHandler) Handle(ctx context.Context, query *GetAccountByIdQuery) (*models.Account, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "getAccountByIdHandler.Handle")
	defer span.Finish()

	if account, err := q.redisRepo.GetAccount(ctx, query.AccountID.String()); err == nil && account != nil {
		return account, nil
	}

	account, err := q.mongoRepo.GetAccountById(ctx, query.AccountID)
	if err != nil {
		return nil, err
	}

	q.redisRepo.PutAccount(ctx, account.AccountID, account)
	return account, nil
}
