package dto

import (
	uuid "github.com/satori/go.uuid"
)

type CreateAccountDto struct {
	AccountID uuid.UUID `json:"account_id" validate:"required"`
	PlayerID  string    `json:"player_id" validate:"required,gte=0,lte=11"`
	Username  string    `json:"username" validate:"required,gte=0,lte=255"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required"`
	IsBan     bool      `json:"is_ban" validate:"boolean"`
}

type CreateAccountResponse struct {
	AccountID uuid.UUID `json:"account_id" validate:"required"`
}

type ChangePwdAccountDto struct {
	AccountID uuid.UUID `json:"account_id" validate:"required"`
	OldPWD    string    `json:"old_password" validate:"required"`
	NewPWD    string    `json:"new_password" validate:"required"`
}

type BanAccountDto struct {
	AccountID uuid.UUID `json:"account_id" validate:"required"`
}

type DeleteAccountDto struct {
	AccountID uuid.UUID `json:"account_id" validate:"required"`
}
