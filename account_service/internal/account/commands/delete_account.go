package commands

import (
	"context"
	"github.com/ce-final-project/backend_rest_api/account_service/config"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/account/repository"
	"github.com/ce-final-project/backend_rest_api/pkg/logger"
	"github.com/opentracing/opentracing-go"
)

type DeleteAccountCmdHandler interface {
	Handle(ctx context.Context, command *DeleteAccountCommand) error
}

type deleteAccountCmdHandler struct {
	log       logger.Logger
	cfg       *config.Config
	mongoRepo repository.Repository
	redisRepo repository.CacheRepository
}

func NewDeleteAccountCmdHandler(log logger.Logger, cfg *config.Config, mongoRepo repository.Repository, redisRepo repository.CacheRepository) *deleteAccountCmdHandler {
	return &deleteAccountCmdHandler{log: log, cfg: cfg, mongoRepo: mongoRepo, redisRepo: redisRepo}
}

func (c *deleteAccountCmdHandler) Handle(ctx context.Context, command *DeleteAccountCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "deleteAccountCmdHandler.Handle")
	defer span.Finish()

	if err := c.mongoRepo.DeleteAccount(ctx, command.AccountID); err != nil {
		return err
	}

	c.redisRepo.DelAccount(ctx, command.AccountID.String())
	return nil
}
