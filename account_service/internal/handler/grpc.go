package handler

import (
	"context"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/core/domain"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/core/ports"
	"github.com/ce-final-project/backend_rest_api/account_service/proto/services"
	"google.golang.org/protobuf/types/known/emptypb"
)

type gRPCHandler struct {
	accSrv ports.AccountService
}

func (g *gRPCHandler) CreateAccount(ctx context.Context, request *services.AccountRequest) (*services.AccountResponse, error) {
	accountPayload := &domain.AccountRequest{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}
	account, err := g.accSrv.CreateAccount(accountPayload)
	if err != nil {
		return nil, err
	}
	return &services.AccountResponse{
		AccountId: account.AccountID,
		PlayerId:  account.PlayerID,
		Username:  account.Username,
		Email:     account.Email,
		IsBan:     account.IsBan,
		CreateAt:  account.CreateAt,
	}, nil
}

func (g *gRPCHandler) GetAllAccount(_ context.Context, _ *emptypb.Empty) (*services.AccountList, error) {
	accounts, err := g.accSrv.GetAllAccount()
	if err != nil {
		return nil, err
	}
	var accountList []*services.AccountResponse
	for _, account := range accounts {
		accountRes := services.AccountResponse{
			AccountId: account.AccountID,
			PlayerId:  account.PlayerID,
			Username:  account.Username,
			Email:     account.Email,
			IsBan:     account.IsBan,
			CreateAt:  account.CreateAt,
		}
		accountList = append(accountList, &accountRes)
	}
	return &services.AccountList{Result: accountList}, nil
}

func (g *gRPCHandler) GetAccount(_ context.Context, message *services.AccountIdMessage) (*services.AccountResponse, error) {
	account, err := g.accSrv.GetAccount(message.AccountId)
	if err != nil {
		return nil, err
	}
	return &services.AccountResponse{
		AccountId: account.AccountID,
		PlayerId:  account.PlayerID,
		Username:  account.Username,
		Email:     account.Email,
		IsBan:     account.IsBan,
		CreateAt:  account.CreateAt,
	}, nil
}

func (g *gRPCHandler) BanAccount(_ context.Context, message *services.AccountIdMessage) (*services.AccountResponse, error) {
	account, err := g.accSrv.BanAccount(message.AccountId)
	if err != nil {
		return nil, err
	}
	return &services.AccountResponse{
		AccountId: account.AccountID,
		PlayerId:  account.PlayerID,
		Username:  account.Username,
		Email:     account.Email,
		IsBan:     account.IsBan,
		CreateAt:  account.CreateAt,
	}, nil

}

func (g *gRPCHandler) ChangePassword(_ context.Context, req *services.ChangePasswordReq) (*services.ChangePasswordRes, error) {
	err := g.accSrv.ChangePassword(req.AccountId, req.OldPassword, req.NewPassword)
	if err != nil {
		return nil, err
	}
	return &services.ChangePasswordRes{Status: "Success"}, nil
}

func NewGRPCHandler(accSrv ports.AccountService) services.AccountServiceServer {
	return &gRPCHandler{accSrv}
}
