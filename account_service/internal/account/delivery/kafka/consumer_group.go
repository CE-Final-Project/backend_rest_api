package kafka

import (
	"context"
	"github.com/ce-final-project/backend_rest_api/account_service/config"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/account/service"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/metrics"
	"github.com/ce-final-project/backend_rest_api/pkg/logger"
	"github.com/go-playground/validator"
	"github.com/segmentio/kafka-go"
	"sync"
)

const (
	PoolSize = 30
)

type accountMessageProcessor struct {
	log     logger.Logger
	cfg     *config.Config
	v       *validator.Validate
	as      *service.AccountService
	metrics *metrics.AccountServiceMetrics
}

func NewAccountMessageProcessor(log logger.Logger, cfg *config.Config, v *validator.Validate, as *service.AccountService, metrics *metrics.AccountServiceMetrics) *accountMessageProcessor {
	return &accountMessageProcessor{log: log, cfg: cfg, v: v, as: as, metrics: metrics}
}

func (s *accountMessageProcessor) ProcessMessages(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		m, err := r.FetchMessage(ctx)
		if err != nil {
			s.log.Warnf("workerID: %v, err: %v", workerID, err)
			continue
		}

		s.logProcessMessage(m, workerID)

		switch m.Topic {
		case s.cfg.KafkaTopics.AccountCreated.TopicName:
			s.processAccountCreated(ctx, r, m)
		case s.cfg.KafkaTopics.AccountUpdated.TopicName:
			s.processAccountUpdated(ctx, r, m)
		case s.cfg.KafkaTopics.AccountDeleted.TopicName:
			s.processAccountDeleted(ctx, r, m)
		}
	}
}
