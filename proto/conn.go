package proto

import (
	"context"
	"io"
	"net"
	"syscall"
)

type Conn struct {
	conn net.Conn
}

func NewConn(conn net.Conn) *Conn {
	return &Conn{
		conn: conn,
	}
}

func (c *Conn) Read(ctx context.Context) (*Header, []byte, error) {
	buf := make([]byte, HeaderLen)
	if _, err := io.ReadFull(c.conn, buf); err != nil {
		return nil, nil, err
	}

	dec := NewDecoder(buf)
	header := new(Header)
	if err := header.DecodeFrom(dec); err != nil {
		return nil, nil, err
	}

	buf = make([]byte, header.BodyLen)
	if _, err := io.ReadFull(c.conn, buf); err != nil {
		return nil, nil, err
	}

	return header, buf, nil
}

func (c *Conn) Write(ctx context.Context, msg Message) error {
	msgEnc := NewEncoder()
	msgEnc.EncodeMessage(msg)
	header := &Header{
		BodyLen: uint64(len(msgEnc.Bytes())),
		Cmd:     uint8(msg.Cmd()),
	}
	headerEnc := NewEncoder()
	headerEnc.EncodeMessage(header)

	buffers := net.Buffers{headerEnc.Bytes(), msgEnc.Bytes()}
	_, err := buffers.WriteTo(c.conn)
	return err
}

func (c *Conn) Call(ctx context.Context, cmd Message, reply Message) error {
	if err := c.Write(ctx, cmd); err != nil {
		return err
	}

	header, buf, err := c.Read(ctx)
	if err != nil {
		return err
	}
	if header.Cmd != uint8(CmdReply) {
		return syscall.EINVAL
	}

	dec := NewDecoder(buf)
	if err := reply.DecodeFrom(dec); err != nil {
		return err
	}

	return nil
}
