package mongo

import (
	"context"
	"log"
	"time"

	"github.com/wenealves10/url-shortener-go/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type dbConnector struct {
	collection *mongo.Collection
}

func (db *dbConnector) Insert(ctx *context.Context, data *model.ShortUrl) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	opts := options.Update().SetUpsert(true)
	_, err := db.collection.UpdateOne(
		*ctx,
		bson.D{{"Hash", data.Hash}},
		bson.D{{"$set", *data}},
		opts,
	)

	return err
}

func (db *dbConnector) FindOne(ctx *context.Context, hashID string) (*model.ShortUrl, error) {
	var result model.ShortUrl
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	err := db.collection.FindOne(*ctx, bson.M{"Hash": hashID}, &options.FindOneOptions{}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetMongoDbConnector(db string, collection string) model.DbConnector {
	return &dbConnector{
		collection: client.Database(db).Collection(collection),
	}
}

func ConnectDB(mongoUri string, timeout time.Duration) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var err error

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoUri)); err != nil {
		log.Fatal(err)
	}

	if err = client.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}
}
