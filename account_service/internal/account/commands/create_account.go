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

type CreateAccountCmdHandler interface {
	Handle(ctx context.Context, command *CreateAccountCommand) error
}

type createAccountHandler struct {
	log       logger.Logger
	cfg       *config.Config
	mongoRepo repository.Repository
	redisRepo repository.CacheRepository
}

func NewCreateAccountHandler(log logger.Logger, cfg *config.Config, mongoRepo repository.Repository, redisRepo repository.CacheRepository) *createAccountHandler {
	return &createAccountHandler{log: log, cfg: cfg, mongoRepo: mongoRepo, redisRepo: redisRepo}
}

func (c *createAccountHandler) Handle(ctx context.Context, command *CreateAccountCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "createAccountHandler.Handle")
	defer span.Finish()

	hashedPWD, err := utils.HashPassword(command.Password)
	if err != nil {
		return err
	}

	account := &models.Account{
		AccountID:    command.AccountID,
		PlayerID:     command.PlayerID,
		Username:     command.Username,
		PasswordHash: hashedPWD,
		Email:        command.Email,
		IsBan:        command.IsBan,
		CreatedAt:    command.CreatedAt,
		UpdatedAt:    command.UpdatedAt,
	}

	created, err := c.mongoRepo.CreateAccount(ctx, account)
	if err != nil {
		return err
	}

	c.redisRepo.PutAccount(ctx, created.AccountID, created)
	return nil
}
