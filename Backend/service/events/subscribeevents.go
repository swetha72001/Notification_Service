package events

import (
	"Backend/dto"
	"Backend/service/sc"
	"context"
	"encoding/json"
	"log"

	natss "github.com/nats-io/nats.go"
)

func SubscribeEvents(ctx context.Context, deps *sc.Dependencies, value chan<- *dto.Events) error {
	natsConn := deps.NatsConn

	sub, err := natsConn.Subscribe("publish-event-topic", func(msg *natss.Msg) {
		log.Println("message ", msg)

		data := &dto.Events{}
		err := json.Unmarshal(msg.Data, data)
		if err != nil {
			log.Println("error while unmarshalling", err)
			return
		}
		value <- data
	})

	go func() {
		<-ctx.Done()
		sub.Drain()
		close(value)
	}()

	return err

}
