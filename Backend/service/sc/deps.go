package sc

import (
	"Backend/config"

	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/mongo"
)

type Dependencies struct {
	MongoClient *mongo.Client
	MongoDb     *mongo.Database
	NsConfig    *config.NSConfig
	NatsConn    *nats.Conn
}
