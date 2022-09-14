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

type CreateAccountCmdHandler interface {
	Handle(ctx context.Context, command *CreateAccountCommand) error
}

type createAccountHandler struct {
	log           logger.Logger
	cfg           *config.Config
	pgRepo        repository.Repository
	kafkaProducer kafkaClient.Producer
}

func NewCreateAccountHandler(log logger.Logger, cfg *config.Config, pgRepo repository.Repository, kafkaProducer kafkaClient.Producer) *createAccountHandler {
	return &createAccountHandler{log: log, cfg: cfg, pgRepo: pgRepo, kafkaProducer: kafkaProducer}
}

func (c *createAccountHandler) Handle(ctx context.Context, command *CreateAccountCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "createAccountHandler.Handle")
	defer span.Finish()

	passwordHashed, err := utils.HashPassword(command.Password)
	if err != nil {
		return err
	}

	accountDto := &models.Account{AccountID: command.AccountID, PlayerID: command.PlayerID, Username: command.Username, Email: command.Email, PasswordHash: passwordHashed}

	account, err := c.pgRepo.CreateAccount(ctx, accountDto)
	if err != nil {
		return err
	}

	msg := &kafkaMessages.AccountCreated{Account: mappers.AccountToGrpcMessage(account)}
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	message := kafka.Message{
		Topic:   c.cfg.KafkaTopics.AccountCreated.TopicName,
		Value:   msgBytes,
		Time:    time.Now().UTC(),
		Headers: tracing.GetKafkaTracingHeadersFromSpanCtx(span.Context()),
	}

	return c.kafkaProducer.PublishMessage(ctx, message)
}
