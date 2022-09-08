package queries

import uuid "github.com/satori/go.uuid"

type AccountQueries struct {
	GetAccountById GetAccountByIdHandler
	SearchAccount  SearchAccountHandler
}

func NewAccountQueries(getAccountById GetAccountByIdHandler, searchAccount SearchAccountHandler) *AccountQueries {
	return &AccountQueries{
		GetAccountById: getAccountById,
		SearchAccount:  searchAccount,
	}
}

type GetAccountByIdQuery struct {
	AccountID uuid.UUID `json:"account_id" validate:"required"`
}

func NewGetAccountByIdQuery(accountId uuid.UUID) *GetAccountByIdQuery {
	return &GetAccountByIdQuery{AccountID: accountId}
}

type SearchAccountQuery struct {
	Username string `json:"username"`
	PlayerID string `json:"player_id"`
}

func NewSearchAccountQuery(username, playerId string) *SearchAccountQuery {
	return &SearchAccountQuery{
		Username: username,
		PlayerID: playerId,
	}
}
