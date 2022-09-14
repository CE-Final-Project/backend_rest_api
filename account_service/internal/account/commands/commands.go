package commands

import (
	uuid "github.com/satori/go.uuid"
)

type AccountCommands struct {
	CreateAccount CreateAccountCmdHandler
	UpdateAccount UpdateAccountCmdHandler
	DeleteAccount DeleteAccountCmdHandler
}

func NewAccountCommands(createAccountHandler CreateAccountCmdHandler, updateAccountHandler UpdateAccountCmdHandler, deleteAccountHandler DeleteAccountCmdHandler) *AccountCommands {
	return &AccountCommands{
		CreateAccount: createAccountHandler,
		UpdateAccount: updateAccountHandler,
		DeleteAccount: deleteAccountHandler,
	}
}

type CreateAccountCommand struct {
	AccountID uuid.UUID `json:"account_id" bson:"_id,omitempty"`
	PlayerID  string    `json:"player_id,omitempty" bson:"player_id,omitempty" validate:"required,max=11"`
	Username  string    `json:"username,omitempty" bson:"username,omitempty" validate:"required,min=3,max=250"`
	Email     string    `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
	Password  string    `json:"password,omitempty" bson:"password" validate:"required,min=8"`
}

func NewCreateAccountCommand(accountID uuid.UUID, playerID string, username string, email string, password string) *CreateAccountCommand {
	return &CreateAccountCommand{
		AccountID: accountID,
		PlayerID:  playerID,
		Username:  username,
		Email:     email,
		Password:  password,
	}
}

type UpdateAccountCommand struct {
	AccountID uuid.UUID `json:"account_id" bson:"_id,omitempty"`
	PlayerID  string    `json:"player_id,omitempty" bson:"player_id,omitempty" validate:"required,max=11"`
	Username  string    `json:"username,omitempty" bson:"username,omitempty" validate:"required,min=3,max=250"`
	Email     string    `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
	Password  string    `json:"password,omitempty" bson:"password" validate:"required,min=8"`
	IsBan     bool      `json:"is_ban,omitempty" bson:"is_ban"`
}

func NewUpdateAccountCommand(accountID uuid.UUID, username string, email string, password string, isBan bool) *UpdateAccountCommand {
	return &UpdateAccountCommand{
		AccountID: accountID,
		Username:  username,
		Email:     email,
		Password:  password,
		IsBan:     isBan,
	}
}

type DeleteAccountCommand struct {
	AccountID uuid.UUID `json:"account_id" bson:"_id,omitempty"`
}

func NewDeleteAccountCommand(accountID uuid.UUID) *DeleteAccountCommand {
	return &DeleteAccountCommand{AccountID: accountID}
}
