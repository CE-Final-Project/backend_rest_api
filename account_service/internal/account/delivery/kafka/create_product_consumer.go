package kafka

import (
	"context"
	"github.com/avast/retry-go"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/account/commands"
	"github.com/ce-final-project/backend_rest_api/pkg/tracing"
	kafkaMessages "github.com/ce-final-project/backend_rest_api/proto/kafka"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"time"
)

const (
	retryAttempts = 3
	retryDelay    = 300 * time.Millisecond
)

var (
	retryOptions = []retry.Option{retry.Attempts(retryAttempts), retry.Delay(retryDelay), retry.DelayType(retry.BackOffDelay)}
)

func (s *accountMessageProcessor) processAccountCreated(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	s.metrics.CreateAccountKafkaMessages.Inc()

	ctx, span := tracing.StartKafkaConsumerTracerSpan(ctx, m.Headers, "accountMessageProcessor.processAccountCreated")
	defer span.Finish()

	msg := &kafkaMessages.AccountCreated{}
	if err := proto.Unmarshal(m.Value, msg); err != nil {
		s.log.WarnMsg("proto.Unmarshal", err)
		s.commitErrMessage(ctx, r, m)
		return
	}

	p := msg.GetAccount()
	command := commands.NewCreateAccountCommand(p.GetAccountID(), p.GetPlayerID(), p.GetUsername(), p.GetEmail(), p.GetPassword(), p.GetIsBan(), p.GetCreatedAt().AsTime(), p.GetUpdatedAt().AsTime())
	if err := s.v.StructCtx(ctx, command); err != nil {
		s.log.WarnMsg("validate", err)
		s.commitErrMessage(ctx, r, m)
		return
	}

	if err := retry.Do(func() error {
		return s.as.Commands.CreateAccount.Handle(ctx, command)
	}, append(retryOptions, retry.Context(ctx))...); err != nil {
		s.log.WarnMsg("CreateAccount.Handle", err)
		s.metrics.ErrorKafkaMessages.Inc()
		return
	}

	s.commitMessage(ctx, r, m)
}
