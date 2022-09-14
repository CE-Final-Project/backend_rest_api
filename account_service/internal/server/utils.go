package server

import (
	"context"
	"github.com/ce-final-project/backend_rest_api/pkg/constants"
	kafkaClient "github.com/ce-final-project/backend_rest_api/pkg/kafka"
	"github.com/heptiolabs/healthcheck"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/segmentio/kafka-go"
	"net"
	"net/http"
	"strconv"
	"time"
)

const (
	stackSize = 1 << 10 // 1 KB
)

func (s *server) connectKafkaBrokers(ctx context.Context) error {
	kafkaConn, err := kafkaClient.NewKafkaConn(ctx, s.cfg.Kafka)
	if err != nil {
		return errors.Wrap(err, "kafka.NewKafkaCon")
	}

	s.kafkaConn = kafkaConn

	brokers, err := kafkaConn.Brokers()
	if err != nil {
		return errors.Wrap(err, "kafkaConn.Brokers")
	}

	s.log.Infof("kafka connected to brokers: %+v", brokers)

	return nil
}

func (s *server) initKafkaTopics(ctx context.Context) {
	controller, err := s.kafkaConn.Controller()
	if err != nil {
		s.log.WarnMsg("kafkaConn.Controller", err)
		return
	}

	controllerURI := net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port))
	s.log.Infof("kafka controller uri: %s", controllerURI)

	conn, err := kafka.DialContext(ctx, "tcp", controllerURI)
	if err != nil {
		s.log.WarnMsg("initKafkaTopics.DialContext", err)
		return
	}
	defer conn.Close() // nolint: errcheck

	s.log.Infof("established new kafka controller connection: %s", controllerURI)

	accountCreateTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.AccountCreate.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.AccountCreate.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.AccountCreate.ReplicationFactor,
	}

	accountCreatedTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.AccountCreated.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.AccountCreated.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.AccountCreated.ReplicationFactor,
	}

	accountUpdateTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.AccountUpdate.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.AccountUpdate.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.AccountUpdate.ReplicationFactor,
	}

	accountUpdatedTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.AccountUpdated.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.AccountUpdated.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.AccountUpdated.ReplicationFactor,
	}

	accountDeleteTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.AccountDelete.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.AccountDelete.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.AccountDelete.ReplicationFactor,
	}

	accountDeletedTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.AccountDeleted.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.AccountDeleted.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.AccountDeleted.ReplicationFactor,
	}

	if err := conn.CreateTopics(
		accountCreateTopic,
		accountUpdateTopic,
		accountCreatedTopic,
		accountUpdatedTopic,
		accountDeleteTopic,
		accountDeletedTopic,
	); err != nil {
		s.log.WarnMsg("kafkaConn.CreateTopics", err)
		return
	}

	s.log.Infof("kafka topics created or already exists: %+v", []kafka.TopicConfig{accountCreateTopic, accountUpdateTopic, accountCreatedTopic, accountUpdatedTopic, accountDeleteTopic, accountDeletedTopic})
}

func (s *server) getConsumerGroupTopics() []string {
	return []string{
		s.cfg.KafkaTopics.AccountCreate.TopicName,
		s.cfg.KafkaTopics.AccountUpdate.TopicName,
		s.cfg.KafkaTopics.AccountDelete.TopicName,
	}
}

func (s *server) runHealthCheck(ctx context.Context) {
	health := healthcheck.NewHandler()

	health.AddLivenessCheck(s.cfg.ServiceName, healthcheck.AsyncWithContext(ctx, func() error {
		return nil
	}, time.Duration(s.cfg.Probes.CheckIntervalSeconds)*time.Second))

	health.AddReadinessCheck(constants.Postgres, healthcheck.AsyncWithContext(ctx, func() error {
		return s.pgConn.Ping(ctx)
	}, time.Duration(s.cfg.Probes.CheckIntervalSeconds)*time.Second))

	health.AddReadinessCheck(constants.Kafka, healthcheck.AsyncWithContext(ctx, func() error {
		_, err := s.kafkaConn.Brokers()
		if err != nil {
			return err
		}
		return nil
	}, time.Duration(s.cfg.Probes.CheckIntervalSeconds)*time.Second))

	go func() {
		s.log.Infof("Account microservice Kubernetes probes listening on port: %s", s.cfg.Probes.Port)
		if err := http.ListenAndServe(s.cfg.Probes.Port, health); err != nil {
			s.log.WarnMsg("ListenAndServe", err)
		}
	}()
}

func (s *server) runMetrics(cancel context.CancelFunc) {
	metricsServer := echo.New()
	go func() {
		metricsServer.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
			StackSize:         stackSize,
			DisablePrintStack: true,
			DisableStackAll:   true,
		}))
		metricsServer.GET(s.cfg.Probes.PrometheusPath, echo.WrapHandler(promhttp.Handler()))
		s.log.Infof("Metrics server is running on port: %s", s.cfg.Probes.PrometheusPort)
		if err := metricsServer.Start(s.cfg.Probes.PrometheusPort); err != nil {
			s.log.Errorf("metricsServer.Start: %v", err)
			cancel()
		}
	}()
}
