package postgres

import (
	"fmt"
	"github.com/ce-final-project/backend_rest_api/pkg/constants"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"os"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	DBName   string `yaml:"dbName"`
	SSLMode  string `yaml:"sslMode"`
	Password string `yaml:"password"`
}

func NewPostgresDB(cfg *Config) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode,
	)

	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitTableDB(db *sqlx.DB) error {
	var initTablePath string
	initTablePathFromEnv := os.Getenv(constants.InitTablePath)
	if initTablePathFromEnv != "" {
		initTablePath = initTablePathFromEnv
	} else {
		getwd, err := os.Getwd()
		if err != nil {
			return errors.Wrap(err, "Initial Table database error: os.Getwd")
		}
		initTablePath = fmt.Sprintf("%s/scripts/account.sql", getwd)
	}
	query, err := os.ReadFile(initTablePath)
	if err != nil {
		return errors.Wrap(err, "Initial Table database error: os.ReadFile")
	}
	_, err = db.Exec(string(query))
	if err != nil {
		return errors.Wrap(err, "Initial Table database error: db.Exec")
	}
	return nil

}
