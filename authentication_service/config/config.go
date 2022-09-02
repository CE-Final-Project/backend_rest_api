package config

import (
	"flag"
	"fmt"
	"github.com/ce-final-project/backend_rest_api/pkg/constants"
	"github.com/ce-final-project/backend_rest_api/pkg/postgres"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "Authentication microservice config path")
}

type Config struct {
	ServiceName string
	Postgresql  *postgres.Config `mapstructure:"postgres"`
	GRPC        GRPC             `mapstructure:"grpc"`
}

type GRPC struct {
	Port        string `mapstructure:"port"`
	Development bool   `mapstructure:"development"`
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
			configPath = fmt.Sprintf("%s/config/config.yaml", getwd)
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
	postgresUser := os.Getenv(constants.PostgresqlUser)
	if postgresUser != "" {
		cfg.Postgresql.User = postgresUser
	}
	postgresPWD := os.Getenv(constants.PostgresqlPassword)
	if postgresPWD != "" {
		cfg.Postgresql.Password = postgresPWD
	}
	postgresDBName := os.Getenv(constants.PostgresqlDBName)
	if postgresDBName != "" {
		cfg.Postgresql.DBName = postgresDBName
	}

	postgresSSL := os.Getenv(constants.PostgresqlSSL)
	if postgresSSL != "" {
		cfg.Postgresql.SSLMode = postgresSSL
	}
	postgresPort := os.Getenv(constants.PostgresqlPort)
	if postgresPort != "" {
		cfg.Postgresql.Port = postgresPort
	}

	return cfg, nil
}
