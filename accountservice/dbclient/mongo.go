package dbclient

import (
	"context"
	"github.com/ce-final-project/backend_rest_api/accountservice/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"strconv"
	"time"
)

type IMongo interface {
	Connect()
	FindAccount(accountId string) (model.Account, error)
	Seed()
}

type Mongo struct {
	DB *mongo.Collection
}

func (m *Mongo) Connect() {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalln("Mongo: Can not connect to DB ===> ", err)
	}
	m.DB = client.Database("GameOnline").Collection("account")
}

func (m *Mongo) FindAccount(accountId string) (model.Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result := model.Account{}
	filter := bson.M{"_id": accountId}
	err := m.DB.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return model.Account{}, err
	}
	return result, nil
}

func (m *Mongo) Seed() {
	var newAccounts []interface{}

	for i := 0; i < 100; i++ {
		account := bson.M{
			"name": "test" + strconv.Itoa(i),
		}
		newAccounts = append(newAccounts, account)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := m.DB.InsertMany(ctx, newAccounts)
	if err != nil {
		log.Printf("Mongo: Can not seed fake account ====> error: %v", err)
		return
	}
}
