package proto

import (
	"context"
	"net"
)

type Conn struct {
	net.Conn
}

func NewConn(conn net.Conn) *Conn {
	return &Conn{
		Conn: conn,
	}
}

func (c *Conn) Read(ctx context.Context, msg Message) error {
	return nil
}

func (c *Conn) Write(ctx context.Context, msg Message) error {
	return nil
}
