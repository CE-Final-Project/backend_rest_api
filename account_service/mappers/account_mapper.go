package mappers

import (
	"github.com/ce-final-project/backend_rest_api/account_service/internal/models"
	accountService "github.com/ce-final-project/backend_rest_api/account_service/proto/account"
	kafkaMessages "github.com/ce-final-project/backend_rest_api/proto/kafka"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func AccountToGrpcMessage(account *models.Account) *kafkaMessages.Account {
	return &kafkaMessages.Account{
		AccountID:      account.AccountID.String(),
		PlayerID:       account.PlayerID,
		Username:       account.Username,
		Email:          account.Email,
		PasswordHashed: account.PasswordHash,
		IsBan:          account.IsBan,
		CreatedAt:      timestamppb.New(account.CreatedAt),
		UpdatedAt:      timestamppb.New(account.UpdatedAt),
	}
}

func AccountFromGrpcMessage(account *kafkaMessages.Account) (*models.Account, error) {

	proUUID, err := uuid.FromString(account.GetAccountID())
	if err != nil {
		return nil, err
	}

	return &models.Account{
		AccountID:    proUUID,
		PlayerID:     account.GetPlayerID(),
		Username:     account.GetUsername(),
		Email:        account.GetEmail(),
		PasswordHash: account.GetPasswordHashed(),
		IsBan:        account.GetIsBan(),
		CreatedAt:    account.GetCreatedAt().AsTime(),
		UpdatedAt:    account.GetUpdatedAt().AsTime(),
	}, nil
}

func AccountToGrpc(account *models.Account) *accountService.Account {
	return &accountService.Account{
		AccountID:      account.AccountID.String(),
		PlayerID:       account.PlayerID,
		Username:       account.Username,
		Email:          account.Email,
		PasswordHashed: account.PasswordHash,
		IsBan:          account.IsBan,
		CreatedAt:      timestamppb.New(account.CreatedAt),
		UpdatedAt:      timestamppb.New(account.UpdatedAt),
	}
}
