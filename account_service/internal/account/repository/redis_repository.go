package repository

import (
	"context"
	"encoding/json"
	"github.com/ce-final-project/backend_rest_api/account_service/config"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/models"
	"github.com/ce-final-project/backend_rest_api/pkg/logger"
	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

const (
	redisAccountPrefixKey = "account"
)

type redisRepository struct {
	log         logger.Logger
	cfg         *config.Config
	redisClient redis.UniversalClient
}

func NewRedisRepository(log logger.Logger, cfg *config.Config, redisClient redis.UniversalClient) *redisRepository {
	return &redisRepository{log: log, cfg: cfg, redisClient: redisClient}
}

func (r *redisRepository) PutAccount(ctx context.Context, key string, account *models.Account) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "redisRepository.PutAccount")
	defer span.Finish()

	accountBytes, err := json.Marshal(account)
	if err != nil {
		r.log.WarnMsg("json.Marshal", err)
		return
	}

	if err := r.redisClient.HSetNX(ctx, r.getRedisAccountPrefixKey(), key, accountBytes).Err(); err != nil {
		r.log.WarnMsg("redisClient.HSetNX", err)
		return
	}
	r.log.Debugf("HSetNX prefix: %s, key: %s", r.getRedisAccountPrefixKey(), key)
}

func (r *redisRepository) GetAccount(ctx context.Context, key string) (*models.Account, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "redisRepository.GetAccount")
	defer span.Finish()

	accountBytes, err := r.redisClient.HGet(ctx, r.getRedisAccountPrefixKey(), key).Bytes()
	if err != nil {
		if err != redis.Nil {
			r.log.WarnMsg("redisClient.HGet", err)
		}
		return nil, errors.Wrap(err, "redisClient.HGet")
	}

	var account models.Account
	if err := json.Unmarshal(accountBytes, &account); err != nil {
		return nil, err
	}

	r.log.Debugf("HGet prefix: %s, key: %s", r.getRedisAccountPrefixKey(), key)
	return &account, nil
}

func (r *redisRepository) DelAccount(ctx context.Context, key string) {
	if err := r.redisClient.HDel(ctx, r.getRedisAccountPrefixKey(), key).Err(); err != nil {
		r.log.WarnMsg("redisClient.HDel", err)
		return
	}
	r.log.Debugf("HDel prefix: %s, key: %s", r.getRedisAccountPrefixKey(), key)
}

func (r *redisRepository) DelAllAccounts(ctx context.Context) {
	if err := r.redisClient.Del(ctx, r.getRedisAccountPrefixKey()).Err(); err != nil {
		r.log.WarnMsg("redisClient.HDel", err)
		return
	}
	r.log.Debugf("Del key: %s", r.getRedisAccountPrefixKey())
}

func (r *redisRepository) getRedisAccountPrefixKey() string {
	if r.cfg.ServiceSettings.RedisAccountPrefixKey != "" {
		return r.cfg.ServiceSettings.RedisAccountPrefixKey
	}

	return redisAccountPrefixKey
}
