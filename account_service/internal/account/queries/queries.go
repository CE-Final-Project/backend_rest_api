package queries

import (
	"github.com/ce-final-project/backend_rest_api/pkg/utils"
	uuid "github.com/satori/go.uuid"
)

type AccountQueries struct {
	GetAccountById GetAccountByIdHandler
	SearchAccount  SearchAccountHandler
}

func NewAccountQueries(getAccountById GetAccountByIdHandler, searchAccount SearchAccountHandler) *AccountQueries {
	return &AccountQueries{GetAccountById: getAccountById, SearchAccount: searchAccount}
}

type GetAccountByIdQuery struct {
	AccountID uuid.UUID `json:"account_id" bson:"_id,omitempty"`
}

func NewGetAccountByIdQuery(accountID uuid.UUID) *GetAccountByIdQuery {
	return &GetAccountByIdQuery{AccountID: accountID}
}

type SearchAccountQuery struct {
	Text       string            `json:"text"`
	Pagination *utils.Pagination `json:"pagination"`
}

func NewSearchAccountQuery(text string, pagination *utils.Pagination) *SearchAccountQuery {
	return &SearchAccountQuery{Text: text, Pagination: pagination}
}
