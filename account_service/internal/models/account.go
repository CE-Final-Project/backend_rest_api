package models

import (
	accountService "github.com/ce-final-project/backend_rest_api/account_service/proto/account"
	"github.com/ce-final-project/backend_rest_api/pkg/utils"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Account struct {
	AccountID    string    `json:"account_id" bson:"_id,omitempty"`
	PlayerID     string    `json:"player_id,omitempty" bson:"player_id,omitempty" validate:"required,max=11"`
	Username     string    `json:"username,omitempty" bson:"username,omitempty" validate:"required,min=3,max=250"`
	Email        string    `json:"email,omitempty" bson:"email,omitempty" validate:"required"`
	PasswordHash string    `json:"password_hash,omitempty" bson:"password_hash,omitempty" validate:"required"`
	IsBan        bool      `json:"is_ban,omitempty" bson:"is_ban,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

// AccountsList accounts list response with pagination
type AccountsList struct {
	TotalCount int64      `json:"totalCount" bson:"totalCount"`
	TotalPages int64      `json:"totalPages" bson:"totalPages"`
	Page       int64      `json:"page" bson:"page"`
	Size       int64      `json:"size" bson:"size"`
	HasMore    bool       `json:"hasMore" bson:"hasMore"`
	Accounts   []*Account `json:"accounts" bson:"accounts"`
}

func NewAccountListWithPagination(accounts []*Account, count int64, pagination *utils.Pagination) *AccountsList {
	return &AccountsList{
		TotalCount: count,
		TotalPages: int64(pagination.GetTotalPages(int(count))),
		Page:       int64(pagination.GetPage()),
		Size:       int64(pagination.GetSize()),
		HasMore:    pagination.GetHasMore(int(count)),
		Accounts:   accounts,
	}
}

func AccountToGrpcMessage(account *Account) *accountService.Account {
	return &accountService.Account{
		AccountID:      account.AccountID,
		PlayerID:       account.PlayerID,
		Username:       account.Username,
		Email:          account.Email,
		PasswordHashed: account.PasswordHash,
		IsBan:          account.IsBan,
		CreatedAt:      timestamppb.New(account.CreatedAt),
		UpdatedAt:      timestamppb.New(account.UpdatedAt),
	}
}

func AccountListToGrpc(accounts *AccountsList) *accountService.SearchRes {
	list := make([]*accountService.Account, 0, len(accounts.Accounts))
	for _, account := range accounts.Accounts {
		list = append(list, AccountToGrpcMessage(account))
	}

	return &accountService.SearchRes{
		TotalCount: accounts.TotalCount,
		TotalPages: accounts.TotalPages,
		Page:       accounts.Page,
		Size:       accounts.Size,
		HasMore:    accounts.HasMore,
		Accounts:   list,
	}
}
