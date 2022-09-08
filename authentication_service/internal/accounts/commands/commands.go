package commands

import "github.com/ce-final-project/backend_rest_api/authentication_service/internal/dto"

type CreateAccountCmdHandler struct {
	CreateDto *dto.CreateAccountDto
}

type UpdateAccountCmdHandler struct {
	UpdateDto *dto.UpdateAccountDto
}

type ChangePasswordCmdHandler struct {
	ChangePwdDto *dto.ChangePwdAccountDto
}

type BanAccountCmdHandler struct {
	BanDto *dto.BanAccountDto
}

type DeleteAccountCmdHandler struct {
	DeleteDto *dto.DeleteAccountDto
}

type AccountCommands struct {
	CreateAccount    CreateAccountCmdHandler
	UpdateAccount    UpdateAccountCmdHandler
	ChangePwdAccount ChangePasswordCmdHandler
	BanAccount       BanAccountCmdHandler
	DeleteAccount    DeleteAccountCmdHandler
}
