package main

import (
	dto "Backend/dto"
	internal "Backend/internal"
	service "Backend/service/events"
	notify "Backend/service/notification"
	deps "Backend/service/sc"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

func main() {
	fmt.Println("Notification Service")

	mainctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	deps := &deps.Dependencies{}
	err := internal.LoadConfig(mainctx, deps)
	if err != nil {
		log.Fatal("error while loading config")
	}
	log.Println("getconfig", deps.NsConfig.GetConfig().MongoDatabaseName)

	mongoErr := internal.InitMongo(mainctx, deps)
	if mongoErr != nil {
		log.Fatal("mongoErr while connecting DB", mongoErr)
	}

	natsErr := internal.NatsInit(mainctx, deps)
	if natsErr != nil {
		log.Println("error while nats connection", natsErr)
	}

	ch := make(chan *dto.Events, 100)
	service.SubscribeEvents(mainctx, deps, ch)

	//registering in dispatcher

	provider := []internal.NotoficationProvider{
		&notify.NatsNotificationService{Conn: deps.NatsConn},
	}
	d := internal.DispatchRegister(provider, deps.MongoDb, deps.NatsConn)
	workers := 5
	var Wg sync.WaitGroup
	internal.Start(mainctx, workers, d, ch, &Wg)

	go func() {
		httpErr := httpInit(ch)
		if httpErr != nil {
			log.Fatal("Err while connecting port", httpErr)
		}
	}()

	Wg.Wait()
	log.Println("signal received sucessfully")
}

func httpInit(ch chan *dto.Events) error {
	hs := internal.HttpSerivice{EventChn: ch}
	mux := http.NewServeMux()
	mux.HandleFunc("/notify", hs.NotifyHandler)

	httpErr := http.ListenAndServe(":8080", mux)
	if httpErr != nil {
		log.Println("Err while connecting port", httpErr)
		return httpErr
	}
	return nil
}
