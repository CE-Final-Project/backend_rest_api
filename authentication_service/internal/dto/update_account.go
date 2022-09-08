package dto

import uuid "github.com/satori/go.uuid"

type UpdateAccountDto struct {
	AccountID uuid.UUID `json:"account_id" validate:"required"`
	Username  string    `json:"username" validate:"required,gte=0,lte=255"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required"`
}
