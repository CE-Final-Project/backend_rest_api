package config

import (
	"flag"
	"fmt"
	"github.com/ce-final-project/backend_rest_api/pkg/constants"
	kafkaClient "github.com/ce-final-project/backend_rest_api/pkg/kafka"
	"github.com/ce-final-project/backend_rest_api/pkg/logger"
	"github.com/ce-final-project/backend_rest_api/pkg/mongodb"
	"github.com/ce-final-project/backend_rest_api/pkg/postgres"
	"github.com/ce-final-project/backend_rest_api/pkg/probes"
	"github.com/ce-final-project/backend_rest_api/pkg/redis"
	"github.com/ce-final-project/backend_rest_api/pkg/tracing"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "Account microservice config path")
}

type Config struct {
	ServiceName      string              `mapstructure:"serviceName"`
	Logger           *logger.Config      `mapstructure:"logger"`
	KafkaTopics      KafkaTopics         `mapstructure:"kafkaTopics"`
	GRPC             GRPC                `mapstructure:"grpc"`
	Postgresql       *postgres.Config    `mapstructure:"postgres"`
	Kafka            *kafkaClient.Config `mapstructure:"kafka"`
	Mongo            *mongodb.Config     `mapstructure:"mongo"`
	Redis            *redis.Config       `mapstructure:"redis"`
	MongoCollections MongoCollections    `mapstructure:"mongoCollections"`
	Probes           probes.Config       `mapstructure:"probes"`
	ServiceSettings  ServiceSettings     `mapstructure:"serviceSettings"`
	Jaeger           *tracing.Config     `mapstructure:"jaeger"`
}

type GRPC struct {
	Port        string `mapstructure:"port"`
	Development bool   `mapstructure:"development"`
}

type MongoCollections struct {
	Accounts string `mapstructure:"accounts"`
}

type KafkaTopics struct {
	AccountCreated kafkaClient.TopicConfig `mapstructure:"accountCreated"`
	AccountUpdated kafkaClient.TopicConfig `mapstructure:"accountUpdated"`
	AccountDeleted kafkaClient.TopicConfig `mapstructure:"accountDeleted"`
}

type ServiceSettings struct {
	RedisAccountPrefixKey string `mapstructure:"redisAccountPrefixKey"`
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

	var cfg *Config
	viper.SetConfigType(constants.Yaml)
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(&cfg); err != nil {
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
