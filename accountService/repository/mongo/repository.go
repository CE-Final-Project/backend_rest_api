package mongo

import (
	"context"
	"fmt"
	"github.com/ce-final-project/backend_rest_api/accountService/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type mongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewMongoRepository(mongoURL string, mongoDB string, mongoTimeout int) (core.AccountRepository, error) {
	repo := &mongoRepository{
		database: mongoDB,
		timeout:  time.Duration(mongoTimeout) * time.Second,
	}
	client, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, fmt.Errorf("%v repository.NewMongoRepository", err)
	}
	repo.client = client
	return repo, nil
}

func (m *mongoRepository) Find(playerId string, accountId string) (*core.Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()
	account := &core.Account{}
	var filter bson.M
	if playerId != "" {
		filter["player_id"] = playerId
	}
	if accountId != "" {
		filter["_id"] = accountId
	}
	collection := m.client.Database(m.database).Collection("accounts")
	err := collection.FindOne(ctx, filter).Decode(&account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%v repository.Account.Find", core.ErrAccountNotFound)
		}
		return nil, fmt.Errorf("%v repository.Account.Find", err)
	}
	return account, nil
}

func (m *mongoRepository) Store(account *core.Account) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()
	collection := m.client.Database(m.database).Collection("accounts")
	_, err := collection.InsertOne(ctx, bson.M{
		"player_id":     account.PlayerId,
		"username":      account.Username,
		"email":         account.Email,
		"password_hash": account.PasswordHash,
		"created_at":    account.CreatedAt,
	})
	if err != nil {
		return fmt.Errorf("%v repository.Account.Store", err)
	}
	return nil
}

func (m *mongoRepository) Remove(playerId string, accountId string) (*core.Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()
	account := &core.Account{}
	var filter bson.M
	if playerId != "" {
		filter["player_id"] = playerId
	}
	if accountId != "" {
		filter["_id"] = accountId
	}
	collection := m.client.Database(m.database).Collection("accounts")
	err := collection.FindOneAndDelete(ctx, filter).Decode(&account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%v repository.Account.Remove", core.ErrAccountNotFound)
		}
		return nil, fmt.Errorf("%v repository.Account.Remove", err)
	}
	return account, nil
}
