package commands

import (
	"context"
	"github.com/ce-final-project/backend_rest_api/account_service/config"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/account/repository"
	kafkaClient "github.com/ce-final-project/backend_rest_api/pkg/kafka"
	"github.com/ce-final-project/backend_rest_api/pkg/logger"
	"github.com/ce-final-project/backend_rest_api/pkg/tracing"
	kafkaMessages "github.com/ce-final-project/backend_rest_api/proto/kafka"
	"github.com/opentracing/opentracing-go"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"time"
)

type DeleteAccountCmdHandler interface {
	Handle(ctx context.Context, command *DeleteAccountCommand) error
}

type deleteAccountHandler struct {
	log           logger.Logger
	cfg           *config.Config
	pgRepo        repository.Repository
	kafkaProducer kafkaClient.Producer
}

func NewDeleteAccountHandler(log logger.Logger, cfg *config.Config, pgRepo repository.Repository, kafkaProducer kafkaClient.Producer) *deleteAccountHandler {
	return &deleteAccountHandler{log: log, cfg: cfg, pgRepo: pgRepo, kafkaProducer: kafkaProducer}
}

func (c *deleteAccountHandler) Handle(ctx context.Context, command *DeleteAccountCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "deleteAccountHandler.Handle")
	defer span.Finish()

	if err := c.pgRepo.DeleteAccountByID(ctx, command.AccountID); err != nil {
		return err
	}

	msg := &kafkaMessages.AccountDeleted{AccountID: command.AccountID.String()}
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	message := kafka.Message{
		Topic:   c.cfg.KafkaTopics.AccountDeleted.TopicName,
		Value:   msgBytes,
		Time:    time.Now().UTC(),
		Headers: tracing.GetKafkaTracingHeadersFromSpanCtx(span.Context()),
	}

	return c.kafkaProducer.PublishMessage(ctx, message)
}
