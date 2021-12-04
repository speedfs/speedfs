package storage

import (
	"testing"

	"github.com/speedfs/speedfs/proto/storage"
	"github.com/speedfs/speedfs/rpc"
)

func TestNewServer(t *testing.T) {
	handler := storage.NewHandler(NewService())
	rpc.NewServer(":8080", handler)
}
