package rpc

import (
	"context"
	"net"
)

type Client struct {
	addr string
}

func NewClient(addr string) *Client {
	return &Client{
		addr: addr,
	}
}

func (c *Client) Call(ctx context.Context, cmd Message, reply Message) error {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	return NewConn(conn).Call(ctx, cmd, reply)
}
