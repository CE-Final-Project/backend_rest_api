package redis

import (
	"fmt"
	"github.com/ce-final-project/backend_rest_api/accountService/core"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"strconv"
)

type redisRepository struct {
	client *redis.Client
}

func newRedisClient(redisURL string) (*redis.Client, error) {
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opts)
	_, err = client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewRedisRepository(redisURL string) (core.AccountRepository, error) {
	repo := &redisRepository{}
	client, err := newRedisClient(redisURL)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewRedisRepository")
	}
	repo.client = client
	return repo, nil
}

func (r *redisRepository) generateKey(playerId string) string {
	return fmt.Sprintf("account:%s", playerId)
}

func (r *redisRepository) Find(username string) ([]*core.Account, error) {
	// TODO: Implement me
	panic("Implement me repository.Find")
}

func (r *redisRepository) FindOne(playerId string) (*core.Account, error) {
	account := &core.Account{}
	key := r.generateKey(playerId)
	result, err := r.client.HGetAll(key).Result()
	if err != nil {
		return nil, errors.Wrap(err, "repository.Account.FindOne")
	}
	if len(result) == 0 {
		return nil, errors.Wrap(core.ErrAccountNotFound, "repository.Account.FindOne")
	}
	var createdAt int64
	createdAt, err = strconv.ParseInt(result["create_at"], 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "repository.Account.FindOne")
	}
	account.PlayerId = result["player_id"]
	account.Username = result["username"]
	account.Email = result["email"]
	account.PasswordHash = result["password_hash"]
	account.CreatedAt = createdAt
	return account, nil
}

func (r *redisRepository) Store(account *core.Account) error {
	key := r.generateKey(account.PlayerId)
	data := map[string]interface{}{
		"player_id":     account.PlayerId,
		"username":      account.Username,
		"email":         account.Email,
		"password_hash": account.PasswordHash,
		"created_at":    account.CreatedAt,
	}
	_, err := r.client.HMSet(key, data).Result()
	if err != nil {
		return errors.Wrap(err, "repository.Account.Store")
	}
	return nil
}

func (r *redisRepository) Remove(playerId string, _ string) (*core.Account, error) {
	key := r.generateKey(playerId)
	account := &core.Account{}
	result, err := r.client.HGetAll(key).Result()
	if err != nil {
		return nil, errors.Wrap(err, "repository.Account.Remove")
	}
	if len(result) == 0 {
		return nil, errors.Wrap(core.ErrAccountNotFound, "repository.Account.Remove")
	}
	_, err = r.client.Del(key).Result()
	if err != nil {
		return nil, errors.Wrap(err, "repository.Account.Remove")
	}

	var createdAt int64
	createdAt, err = strconv.ParseInt(result["create_at"], 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "repository.Account.Remove")
	}

	account.PlayerId = result["player_id"]
	account.Username = result["username"]
	account.Email = result["email"]
	account.CreatedAt = createdAt
	return account, nil
}
