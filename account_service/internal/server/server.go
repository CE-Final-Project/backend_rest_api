package server

import (
	"context"
	"github.com/ce-final-project/backend_rest_api/account_service/config"
	accountKafka "github.com/ce-final-project/backend_rest_api/account_service/internal/account/delivery/kafka"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/account/repository"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/account/service"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/metrics"
	"github.com/ce-final-project/backend_rest_api/pkg/interceptors"
	kafkaClient "github.com/ce-final-project/backend_rest_api/pkg/kafka"
	"github.com/ce-final-project/backend_rest_api/pkg/logger"
	"github.com/ce-final-project/backend_rest_api/pkg/mongodb"
	redisClient "github.com/ce-final-project/backend_rest_api/pkg/redis"
	"github.com/ce-final-project/backend_rest_api/pkg/tracing"
	"github.com/go-playground/validator"
	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"os/signal"
	"syscall"
)

type server struct {
	log         logger.Logger
	cfg         *config.Config
	v           *validator.Validate
	kafkaConn   *kafka.Conn
	im          interceptors.InterceptorManager
	mongoClient *mongo.Client
	redisClient redis.UniversalClient
	as          *service.AccountService
	metrics     *metrics.AccountServiceMetrics
}

func NewServer(log logger.Logger, cfg *config.Config) *server {
	return &server{log: log, cfg: cfg, v: validator.New()}
}

func (s *server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	s.im = interceptors.NewInterceptorManager(s.log)
	s.metrics = metrics.NewAccountServiceMetrics(s.cfg)

	mongoDBConn, err := mongodb.NewMongoDBConn(ctx, s.cfg.Mongo)
	if err != nil {
		return errors.Wrap(err, "NewMongoDBConn")
	}
	s.mongoClient = mongoDBConn
	defer mongoDBConn.Disconnect(ctx) // nolint: errcheck
	s.log.Infof("Mongo connected: %v", mongoDBConn.NumberSessionsInProgress())

	s.redisClient = redisClient.NewUniversalRedisClient(s.cfg.Redis)
	defer s.redisClient.Close() // nolint: errcheck
	s.log.Infof("Redis connected: %+v", s.redisClient.PoolStats())

	mongoRepo := repository.NewMongoRepository(s.log, s.cfg, s.mongoClient)
	redisRepo := repository.NewRedisRepository(s.log, s.cfg, s.redisClient)

	s.as = service.NewAccountService(s.log, s.cfg, mongoRepo, redisRepo)

	accountMessageProcessor := accountKafka.NewAccountMessageProcessor(s.log, s.cfg, s.v, s.as, s.metrics)

	s.log.Info("Starting Account Kafka consumers")
	cg := kafkaClient.NewConsumerGroup(s.cfg.Kafka.Brokers, s.cfg.Kafka.GroupID, s.log)
	go cg.ConsumeTopic(ctx, s.getConsumerGroupTopics(), accountKafka.PoolSize, accountMessageProcessor.ProcessMessages)

	if err := s.connectKafkaBrokers(ctx); err != nil {
		return errors.Wrap(err, "s.connectKafkaBrokers")
	}
	defer s.kafkaConn.Close() // nolint: errcheck

	s.runHealthCheck(ctx)
	s.runMetrics(cancel)

	if s.cfg.Jaeger.Enable {
		tracer, closer, err := tracing.NewJaegerTracer(s.cfg.Jaeger)
		if err != nil {
			return err
		}
		defer closer.Close() // nolint: errcheck
		opentracing.SetGlobalTracer(tracer)
	}

	closeGrpcServer, grpcServer, err := s.newAccountGrpcServer()
	if err != nil {
		return errors.Wrap(err, "NewScmGrpcServer")
	}
	defer closeGrpcServer() // nolint: errcheck

	<-ctx.Done()
	grpcServer.GracefulStop()
	return nil
}
