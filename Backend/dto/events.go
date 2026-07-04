package dto

import "time"

type Events struct {
	EventID         string
	SubscriberName  string
	Chanel          string
	TransactionType string
	CreatedDate     time.Time
	Amount          string
}
