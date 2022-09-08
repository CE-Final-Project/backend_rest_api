package service

import "github.com/ce-final-project/backend_rest_api/authentication_service/internal/accounts/commands"

type AccountService struct {
	Commands *commands.AccountCommands
	Queries  *queries.ProductQueries
}
