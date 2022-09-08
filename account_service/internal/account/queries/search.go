package queries

import (
	"context"
	"github.com/ce-final-project/backend_rest_api/account_service/config"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/account/repository"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/models"
	"github.com/ce-final-project/backend_rest_api/pkg/logger"
)

type SearchAccountHandler interface {
	Handle(ctx context.Context, query *SearchAccountQuery) (*models.AccountsList, error)
}

type searchAccountHandler struct {
	log       logger.Logger
	cfg       *config.Config
	mongoRepo repository.Repository
	redisRepo repository.CacheRepository
}

func NewSearchAccountHandler(log logger.Logger, cfg *config.Config, mongoRepo repository.Repository, redisRepo repository.CacheRepository) *searchAccountHandler {
	return &searchAccountHandler{log: log, cfg: cfg, mongoRepo: mongoRepo, redisRepo: redisRepo}
}

func (s *searchAccountHandler) Handle(ctx context.Context, query *SearchAccountQuery) (*models.AccountsList, error) {
	return s.mongoRepo.Search(ctx, query.Text, query.Pagination)
}
