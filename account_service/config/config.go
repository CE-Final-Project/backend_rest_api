package config

import (
	"fmt"
	"github.com/ce-final-project/backend_rest_api/pkg/constants"
	kafkaClient "github.com/ce-final-project/backend_rest_api/pkg/kafka"
	"github.com/ce-final-project/backend_rest_api/pkg/logger"
	"github.com/ce-final-project/backend_rest_api/pkg/postgres"
	"github.com/ce-final-project/backend_rest_api/pkg/probes"
	"github.com/ce-final-project/backend_rest_api/pkg/tracing"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
)

var configPath string

type Config struct {
	ServiceName string              `mapstructure:"serviceName"`
	Logger      *logger.Config      `mapstructure:"logger"`
	KafkaTopics KafkaTopics         `mapstructure:"kafkaTopics"`
	GRPC        GRPC                `mapstructure:"grpc"`
	Postgresql  *postgres.Config    `mapstructure:"postgres"`
	Kafka       *kafkaClient.Config `mapstructure:"kafka"`
	Probes      probes.Config       `mapstructure:"probes"`
	Jaeger      *tracing.Config     `mapstructure:"jaeger"`
}

type GRPC struct {
	Port        string `mapstructure:"port"`
	Development bool   `mapstructure:"development"`
}

type KafkaTopics struct {
	AccountCreate  kafkaClient.TopicConfig `mapstructure:"accountCreate"`
	AccountCreated kafkaClient.TopicConfig `mapstructure:"accountCreated"`
	AccountUpdate  kafkaClient.TopicConfig `mapstructure:"accountUpdate"`
	AccountUpdated kafkaClient.TopicConfig `mapstructure:"accountUpdated"`
	AccountDelete  kafkaClient.TopicConfig `mapstructure:"accountDelete"`
	AccountDeleted kafkaClient.TopicConfig `mapstructure:"accountDeleted"`
}

func InitConfig() (*Config, error) {
	if configPath == "" {
		configPathFromEnv := os.Getenv(constants.ConfigPath)
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			getwd, err := os.Getwd()
			if err != nil {
				return nil, errors.Wrap(err, "os.Getwd")
			}
			configPath = fmt.Sprintf("%s/account_service/config/config.yaml", getwd)
		}
	}

	cfg := &Config{}

	viper.SetConfigType(constants.Yaml)
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshal")
	}

	grpcPort := os.Getenv(constants.GrpcPort)
	if grpcPort != "" {
		cfg.GRPC.Port = grpcPort
	}

	postgresHost := os.Getenv(constants.PostgresqlHost)
	if postgresHost != "" {
		cfg.Postgresql.Host = postgresHost
	}
	postgresPort := os.Getenv(constants.PostgresqlPort)
	if postgresPort != "" {
		cfg.Postgresql.Port = postgresPort
	}
	jaegerAddr := os.Getenv(constants.JaegerHostPort)
	if jaegerAddr != "" {
		cfg.Jaeger.HostPort = jaegerAddr
	}
	kafkaBrokers := os.Getenv(constants.KafkaBrokers)
	if kafkaBrokers != "" {
		cfg.Kafka.Brokers = []string{kafkaBrokers}
	}

	return cfg, nil
}
