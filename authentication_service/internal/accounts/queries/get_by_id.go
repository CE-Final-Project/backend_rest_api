package queries

import (
	"context"
	"github.com/ce-final-project/backend_rest_api/authentication_service/config"
	"github.com/ce-final-project/backend_rest_api/authentication_service/internal/dto"
	"github.com/ce-final-project/backend_rest_api/pkg/logger"
	"github.com/ce-final-project/backend_rest_api/pkg/tracing"
	"github.com/opentracing/opentracing-go"
)

type GetAccountByIdHandler interface {
	Handle(ctx context.Context, query *GetAccountByIdQuery) (*dto.AccountResponse, error)
}

type getAccountByIdHandler struct {
	log      logger.Logger
	cfg      *config.Config
	asClient accountService.AccountServiceClient
}

func NewGetAccountByIdHandler(log logger.Logger, cfg *config.Config, asClient accountService.AccountServiceClient) *getAccountByIdHandler {
	return &getAccountByIdHandler{
		log,
		cfg,
		asClient,
	}
}

func (q *getAccountByIdHandler) Handler(ctx context.Context, query *GetAccountByIdQuery) (*dto.AccountResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "getAccountByIdHandler.Handle")
	defer span.Finish()

	ctx = tracing.InjectTextMapCarrierToGrpcMetaData(ctx, span.Context())
	res, err := q.asClient.GetAccountById(ctx, &accountService.GetAccountIdReq{AccountID: query.AccountID.String()})
	if err != nil {
		return nil, err
	}

	return dto.AccountResponseFromGrpc(res.GetAccount()), nil
}
