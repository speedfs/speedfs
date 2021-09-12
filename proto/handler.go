package proto

import (
	"context"
)

type Handler interface {
	Handle(ctx context.Context, cmd uint8, buf []byte) (Message, error)
}
