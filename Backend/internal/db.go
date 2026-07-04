package internal

import (
	"Backend/service/sc"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitMongo(ctx context.Context, deps *sc.Dependencies) error {
	clientOptions := options.Client().ApplyURI(deps.NsConfig.GetConfig().MongoURL)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println("error while connecting mongo", err)
		return err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println("error while connecting mongo", err)
		return err
	}

	deps.MongoDb = client.Database(deps.NsConfig.MongoDatabaseName)
	deps.MongoClient = client
	return nil
}
