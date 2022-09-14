package service

import (
	"github.com/ce-final-project/backend_rest_api/account_service/config"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/account/commands"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/account/queries"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/account/repository"
	kafkaClient "github.com/ce-final-project/backend_rest_api/pkg/kafka"
	"github.com/ce-final-project/backend_rest_api/pkg/logger"
)

type AccountService struct {
	Commands *commands.AccountCommands
	Queries  *queries.AccountQueries
}

func NewAccountService(log logger.Logger, cfg *config.Config, pgRepo repository.Repository, kafkaProducer kafkaClient.Producer) *AccountService {

	updateAccountHandler := commands.NewUpdateAccountHandler(log, cfg, pgRepo, kafkaProducer)
	createAccountHandler := commands.NewCreateAccountHandler(log, cfg, pgRepo, kafkaProducer)
	deleteAccountHandler := commands.NewDeleteAccountHandler(log, cfg, pgRepo, kafkaProducer)

	getAccountByIdHandler := queries.NewGetAccountByIdHandler(log, cfg, pgRepo)

	accountCommands := commands.NewAccountCommands(createAccountHandler, updateAccountHandler, deleteAccountHandler)
	accountQueries := queries.NewAccountQueries(getAccountByIdHandler)

	return &AccountService{Commands: accountCommands, Queries: accountQueries}
}
