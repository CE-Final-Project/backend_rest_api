package dto

import "time"

type AccountResponse struct {
	AccountID string    `json:"account_id"`
	PlayerID  string    `json:"player_id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func AccountResponseFromGrpc(account *AccountService.account) *AccountResponse {
	return &AccountResponse{
		AccountID: account.GetAccountID(),
		PlayerID:  account.GetPlayerID(),
		Username:  account.GetUsername(),
		Email:     account.GetEmail(),
		CreatedAt: account.GetCreatedAt().AsTime(),
		UpdatedAt: account.GetUpdatedAt().AsTime(),
	}
}
