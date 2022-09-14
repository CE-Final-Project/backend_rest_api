package commands

import (
	"context"
	"github.com/ce-final-project/backend_rest_api/account_service/config"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/account/repository"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/models"
	"github.com/ce-final-project/backend_rest_api/account_service/mappers"
	kafkaClient "github.com/ce-final-project/backend_rest_api/pkg/kafka"
	"github.com/ce-final-project/backend_rest_api/pkg/logger"
	"github.com/ce-final-project/backend_rest_api/pkg/tracing"
	"github.com/ce-final-project/backend_rest_api/pkg/utils"
	kafkaMessages "github.com/ce-final-project/backend_rest_api/proto/kafka"
	"github.com/opentracing/opentracing-go"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"time"
)

type UpdateAccountCmdHandler interface {
	Handle(ctx context.Context, command *UpdateAccountCommand) error
}

type updateAccountHandler struct {
	log           logger.Logger
	cfg           *config.Config
	pgRepo        repository.Repository
	kafkaProducer kafkaClient.Producer
}

func NewUpdateAccountHandler(log logger.Logger, cfg *config.Config, pgRepo repository.Repository, kafkaProducer kafkaClient.Producer) *updateAccountHandler {
	return &updateAccountHandler{log: log, cfg: cfg, pgRepo: pgRepo, kafkaProducer: kafkaProducer}
}

func (c *updateAccountHandler) Handle(ctx context.Context, command *UpdateAccountCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "updateAccountHandler.Handle")
	defer span.Finish()

	passwordHashed, err := utils.HashPassword(command.Password)
	if err != nil {
		return err
	}

	productDto := &models.Account{AccountID: command.AccountID, PlayerID: command.PlayerID, Username: command.Username, Email: command.Email, PasswordHash: passwordHashed, IsBan: command.IsBan}

	product, err := c.pgRepo.UpdateAccount(ctx, productDto)
	if err != nil {
		return err
	}

	msg := &kafkaMessages.AccountUpdated{Account: mappers.AccountToGrpcMessage(product)}
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	message := kafka.Message{
		Topic:   c.cfg.KafkaTopics.AccountUpdated.TopicName,
		Value:   msgBytes,
		Time:    time.Now().UTC(),
		Headers: tracing.GetKafkaTracingHeadersFromSpanCtx(span.Context()),
	}

	return c.kafkaProducer.PublishMessage(ctx, message)
}
