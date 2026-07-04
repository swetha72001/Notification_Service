package internal

import (
	"Backend/service/sc"
	"context"
	"log"
	"time"

	nats "github.com/nats-io/nats.go"
)

func NatsInit(ctx context.Context, deps *sc.Dependencies) error {
	natsoption := nats.RetryOnFailedConnect(true)
	waitoption := nats.ReconnectWait(time.Second * 3)
	conn, err := nats.Connect("localhost:4222", natsoption, waitoption)
	if err != nil || conn == nil {
		log.Println("error while connecting nats", err)
		return err
	}
	if conn != nil {
		deps.NatsConn = conn
		log.Println(" connected nats")
	}
	return nil
}
