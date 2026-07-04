package internal

import (
	"Backend/dto"
	"context"
)

type NotoficationProvider interface {
	Name() string
	SendNotification(ctx context.Context, payload *dto.Payload) error
}
