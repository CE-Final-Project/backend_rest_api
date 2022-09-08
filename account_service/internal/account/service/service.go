package service

import (
	"github.com/ce-final-project/backend_rest_api/account_service/config"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/account/commands"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/account/queries"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/account/repository"
	"github.com/ce-final-project/backend_rest_api/pkg/logger"
)

type AccountService struct {
	Commands *commands.AccountCommands
	Queries  *queries.AccountQueries
}

func NewAccountService(
	log logger.Logger,
	cfg *config.Config,
	mongoRepo repository.Repository,
	redisRepo repository.CacheRepository,
) *AccountService {

	createAccountHandler := commands.NewCreateAccountHandler(log, cfg, mongoRepo, redisRepo)
	deleteAccountCmdHandler := commands.NewDeleteAccountCmdHandler(log, cfg, mongoRepo, redisRepo)
	updateAccountCmdHandler := commands.NewUpdateAccountCmdHandler(log, cfg, mongoRepo, redisRepo)

	getAccountByIdHandler := queries.NewGetAccountByIdHandler(log, cfg, mongoRepo, redisRepo)
	searchAccountHandler := queries.NewSearchAccountHandler(log, cfg, mongoRepo, redisRepo)

	accountCommands := commands.NewAccountCommands(createAccountHandler, updateAccountCmdHandler, deleteAccountCmdHandler)
	accountQueries := queries.NewAccountQueries(getAccountByIdHandler, searchAccountHandler)

	return &AccountService{Commands: accountCommands, Queries: accountQueries}
}
