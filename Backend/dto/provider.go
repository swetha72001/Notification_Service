package dto

import "time"

type Payload struct {
	TransactionType string
	SentDate        time.Time
	CustomerName    string
	Amount          string
}
