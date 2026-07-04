package notification

import (
	"Backend/dto"
	"context"
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

type NatsNotificationService struct {
	Conn *nats.Conn
}

func (n *NatsNotificationService) Name() string {
	return "nats"
}

func (n *NatsNotificationService) SendNotification(ctx context.Context, payload *dto.Payload) error {
	//publish
	req, err := json.Marshal(payload)
	if err != nil {
		log.Println("error while unmarshalling", err)
		return err
	}
	n.Conn.Publish("notify-customer-event", req)
	return nil
}
