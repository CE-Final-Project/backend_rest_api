package repository

import (
	"context"
	"github.com/ce-final-project/backend_rest_api/account_service/config"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/models"
	"github.com/ce-final-project/backend_rest_api/pkg/logger"
	"github.com/ce-final-project/backend_rest_api/pkg/utils"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct {
	log logger.Logger
	cfg *config.Config
	db  *mongo.Client
}

func NewMongoRepository(log logger.Logger, cfg *config.Config, db *mongo.Client) *mongoRepository {
	return &mongoRepository{log: log, cfg: cfg, db: db}
}

func (p *mongoRepository) CreateAccount(ctx context.Context, account *models.Account) (*models.Account, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.CreateAccount")
	defer span.Finish()

	collection := p.db.Database(p.cfg.Mongo.Db).Collection(p.cfg.MongoCollections.Accounts)

	_, err := collection.InsertOne(ctx, account, &options.InsertOneOptions{})
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "InsertOne")
	}

	return account, nil
}

func (p *mongoRepository) UpdateAccount(ctx context.Context, account *models.Account) (*models.Account, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.UpdateAccount")
	defer span.Finish()

	collection := p.db.Database(p.cfg.Mongo.Db).Collection(p.cfg.MongoCollections.Accounts)

	ops := options.FindOneAndUpdate()
	ops.SetReturnDocument(options.After)
	ops.SetUpsert(true)

	var updated models.Account
	if err := collection.FindOneAndUpdate(ctx, bson.M{"_id": account.AccountID}, bson.M{"$set": account}, ops).Decode(&updated); err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "Decode")
	}

	return &updated, nil
}

func (p *mongoRepository) GetAccountById(ctx context.Context, uuid uuid.UUID) (*models.Account, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.GetAccountById")
	defer span.Finish()

	collection := p.db.Database(p.cfg.Mongo.Db).Collection(p.cfg.MongoCollections.Accounts)

	var account models.Account
	if err := collection.FindOne(ctx, bson.M{"_id": uuid.String()}).Decode(&account); err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "Decode")
	}

	return &account, nil
}

func (p *mongoRepository) DeleteAccount(ctx context.Context, uuid uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.DeleteAccount")
	defer span.Finish()

	collection := p.db.Database(p.cfg.Mongo.Db).Collection(p.cfg.MongoCollections.Accounts)

	return collection.FindOneAndDelete(ctx, bson.M{"_id": uuid.String()}).Err()
}

func (p *mongoRepository) Search(ctx context.Context, search string, pagination *utils.Pagination) (*models.AccountsList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.Search")
	defer span.Finish()

	collection := p.db.Database(p.cfg.Mongo.Db).Collection(p.cfg.MongoCollections.Accounts)

	filter := bson.D{
		{Key: "$or", Value: bson.A{
			bson.D{{Key: "username", Value: primitive.Regex{Pattern: search, Options: "gi"}}},
			bson.D{{Key: "player_id", Value: primitive.Regex{Pattern: search, Options: "gi"}}},
		}},
	}

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "CountDocuments")
	}
	if count == 0 {
		return &models.AccountsList{Accounts: make([]*models.Account, 0)}, nil
	}

	limit := int64(pagination.GetLimit())
	skip := int64(pagination.GetOffset())
	cursor, err := collection.Find(ctx, filter, &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
	})
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "Find")
	}
	defer cursor.Close(ctx) // nolint: errcheck

	accounts := make([]*models.Account, 0, pagination.GetSize())

	for cursor.Next(ctx) {
		var prod models.Account
		if err := cursor.Decode(&prod); err != nil {
			p.traceErr(span, err)
			return nil, errors.Wrap(err, "Find")
		}
		accounts = append(accounts, &prod)
	}

	if err := cursor.Err(); err != nil {
		span.SetTag("error", true)
		span.LogKV("error_code", err.Error())
		return nil, errors.Wrap(err, "cursor.Err")
	}

	return models.NewAccountListWithPagination(accounts, count, pagination), nil
}

func (p *mongoRepository) traceErr(span opentracing.Span, err error) {
	span.SetTag("error", true)
	span.LogKV("error_code", err.Error())
}
