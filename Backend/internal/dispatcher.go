package internal

import (
	"Backend/dto"
	"context"
	"log"
	"time"

	nats "github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/mongo"
)

type DispatcherService struct {
	Providers []NotoficationProvider
	MongoDb   *mongo.Database
	NatsConn  *nats.Conn
}

func DispatchRegister(providers []NotoficationProvider, db *mongo.Database, natsConn *nats.Conn) *DispatcherService {
	return &DispatcherService{Providers: providers, MongoDb: db, NatsConn: natsConn}
}

func (d *DispatcherService) DispatcherProcess(ctx context.Context, event *dto.Events) error {

	_, err := d.MongoDb.Collection("Event").InsertOne(ctx, event)
	if err != nil {
		log.Println("error while inserting event", err)
		return err
	}
	for _, d := range d.Providers {

		payload := &dto.Payload{
			TransactionType: event.TransactionType,
			SentDate:        time.Now(),
			Amount:          event.Amount,
		}

		d.SendNotification(ctx, payload)
	}
	return nil
}
