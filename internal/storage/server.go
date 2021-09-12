package storage

import (
	"context"
	"log"
	"net"

	"github.com/speedfs/speedfs/proto"
)

type Server struct {
	addr    string
	handler proto.Handler
}

// NewServer return a new Server
func NewServer(addr string, handler proto.Handler) *Server {
	return &Server{
		addr:    addr,
		handler: handler,
	}
}

func (srv *Server) serve(l net.Listener) error {
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("rpc: accept failed, err: %v", err)
			return err
		}

		go srv.handle(NewConn(conn))
	}
}

func (srv *Server) ListenAndServe() error {
	l, err := net.Listen("tcp", srv.addr)
	if err != nil {
		return err
	}
	return srv.serve(l)
}

func (srv *Server) handle(conn *Conn) error {
	for {
		cmd, buf, err := conn.Read()
		if err != nil {
			return err
		}

		reply, err := srv.handler.Handle(context.TODO(), cmd, buf)
		if err != nil {
			return err
		}

		err = conn.Write(reply)
		if err != nil {
			return err
		}
	}
}
