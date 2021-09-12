package storage

import (
	"net"

	"github.com/speedfs/speedfs/proto"
)

type Conn struct {
	net.Conn
}

func NewConn(conn net.Conn) *Conn {
	return &Conn{
		Conn: conn,
	}
}

func (conn *Conn) Read() (uint8, []byte, error) {
	return 0, nil, nil
}

func (conn *Conn) Write(msg proto.Message) error {
	return nil
}
