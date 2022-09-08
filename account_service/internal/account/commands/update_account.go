package commands

import (
	"context"
	"github.com/ce-final-project/backend_rest_api/account_service/config"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/account/repository"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/models"
	"github.com/ce-final-project/backend_rest_api/pkg/logger"
	"github.com/ce-final-project/backend_rest_api/pkg/utils"
	"github.com/opentracing/opentracing-go"
)

type UpdateAccountCmdHandler interface {
	Handle(ctx context.Context, command *UpdateAccountCommand) error
}

type updateAccountCmdHandler struct {
	log       logger.Logger
	cfg       *config.Config
	mongoRepo repository.Repository
	redisRepo repository.CacheRepository
}

func NewUpdateAccountCmdHandler(log logger.Logger, cfg *config.Config, mongoRepo repository.Repository, redisRepo repository.CacheRepository) *updateAccountCmdHandler {
	return &updateAccountCmdHandler{log: log, cfg: cfg, mongoRepo: mongoRepo, redisRepo: redisRepo}
}

func (c *updateAccountCmdHandler) Handle(ctx context.Context, command *UpdateAccountCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "updateAccountCmdHandler.Handle")
	defer span.Finish()

	hashedPWD, err := utils.HashPassword(command.Password)
	if err != nil {
		return err
	}

	account := &models.Account{
		AccountID:    command.AccountID,
		Username:     command.Username,
		Email:        command.Email,
		PasswordHash: hashedPWD,
		IsBan:        command.IsBan,
		UpdatedAt:    command.UpdatedAt,
	}

	updated, err := c.mongoRepo.UpdateAccount(ctx, account)
	if err != nil {
		return err
	}

	c.redisRepo.PutAccount(ctx, updated.AccountID, updated)
	return nil
}
