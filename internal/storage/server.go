package storage

import (
	"context"
	"log"
	"net"

	"github.com/speedfs/speedfs/rpc"
)

type Server struct {
	addr    string
	handler rpc.Handler
}

// NewServer return a new Server
func NewServer(addr string, handler rpc.Handler) *Server {
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

		go srv.handle(rpc.NewConn(conn))
	}
}

func (srv *Server) ListenAndServe() error {
	l, err := net.Listen("tcp", srv.addr)
	if err != nil {
		return err
	}
	return srv.serve(l)
}

func (srv *Server) handle(conn *rpc.Conn) error {
	for {
		ctx := context.Background()
		header, buf, err := conn.Read(ctx)
		if err != nil {
			return err
		}

		reply, err := srv.handler.Handle(ctx, uint8(header.Cmd), buf)
		if err != nil {
			return err
		}

		err = conn.Write(ctx, reply)
		if err != nil {
			return err
		}
	}
}
