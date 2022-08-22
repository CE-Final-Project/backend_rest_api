package mongo

import (
	"context"
	"fmt"
	"github.com/ce-final-project/backend_rest_api/accountService/core"
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

// TODO: Implement all this method
func (m mongoRepository) Find(playerId string, accountId string) (*core.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (m mongoRepository) Store(account *core.Account) error {
	//TODO implement me
	panic("implement me")
}

func (m mongoRepository) Remove(playerId string, accountId string) (*core.Account, error) {
	//TODO implement me
	panic("implement me")
}
