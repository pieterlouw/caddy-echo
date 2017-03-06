package echoserver

import (
	"fmt"
	"io"
	"net"
)

// Server is a TCP echo implementation of the
// caddy.Server interface type
type Server struct {
	Port     string
	listener net.Listener
}

// NewServer returns a new tcpecho server
func NewServer(p string) (*Server, error) {
	return &Server{
		Port: p,
	}, nil
}

// Listen starts listening by creating a new listener
// and returning it. It does not start accepting
// connections.
func (s *Server) Listen() (net.Listener, error) {
	return net.Listen("tcp", fmt.Sprintf(".:%s", s.Port))
}

// Serve starts serving using the provided listener.
// Serve blocks indefinitely, or in other
// words, until the server is stopped.
func (s *Server) Serve(ln net.Listener) error {
	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}

		go func(c net.Conn) {
			// Echo all incoming data.
			io.Copy(c, c)
			// Shut down the connection.
			c.Close()
		}(conn)
	}
}

// ListenPacket is a no-op to satisfy caddy.Server interface
func (s *Server) ListenPacket() (net.PacketConn, error) { return nil, nil }

// ServePacket is a no-op to satisfy caddy.Server interface
func (s *Server) ServePacket(net.PacketConn) error { return nil }

// OnStartupComplete lists the sites served by this server
// and any relevant information
func (s *Server) OnStartupComplete() {
	fmt.Println("OnStartupComplete:", s.Port)
}
