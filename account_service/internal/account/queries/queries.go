package queries

import (
	uuid "github.com/satori/go.uuid"
)

type AccountQueries struct {
	GetAccountById GetAccountByIdHandler
}

func NewAccountQueries(getAccountById GetAccountByIdHandler) *AccountQueries {
	return &AccountQueries{GetAccountById: getAccountById}
}

type GetAccountByIdQuery struct {
	AccountID uuid.UUID `json:"account_id" validate:"required,gte=0,lte=255"`
}

func NewGetAccountByIdQuery(productID uuid.UUID) *GetAccountByIdQuery {
	return &GetAccountByIdQuery{AccountID: productID}
}
