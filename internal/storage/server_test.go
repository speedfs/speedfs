package storage

import (
	"testing"

	"github.com/speedfs/speedfs/proto/storage"
)

func TestNewServer(t *testing.T) {
	handler := storage.NewHandler(NewService())
	NewServer(":8080", handler)
}
