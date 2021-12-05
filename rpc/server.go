package rpc

import (
	"context"
	"log"
	"net"
)

type Server struct {
	addr    string
	handler Handler
}

// NewServer return a new Server
func NewServer(addr string, handler Handler) *Server {
	return &Server{
		addr:    addr,
		handler: handler,
	}
}

func (s *Server) serve(l net.Listener) error {
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("rpc: accept failed, err: %v", err)
			return err
		}

		go s.handle(NewConn(conn))
	}
}

func (s *Server) ListenAndServe() error {
	l, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	return s.serve(l)
}

func (s *Server) handle(conn *Conn) error {
	for {
		ctx := context.Background()
		header, buf, err := conn.Read(ctx)
		if err != nil {
			return err
		}

		reply, err := s.handler.Handle(ctx, header.Cmd, buf)
		if err != nil {
			return err
		}

		err = conn.Write(ctx, reply)
		if err != nil {
			return err
		}
	}
}
