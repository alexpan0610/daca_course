package p2p

import (
	"net"
)

type Server struct {
	addr string
}

func (s *Server) Start() error {
	listen, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		p := NewPeer(conn)
		p.Start()
		p.Send("I am an inbound peer\n")
	}
}

func (s *Server) Connect(addrs []string) {
	for _, addr := range addrs {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			continue
		}

		p := NewPeer(conn)
		p.Start()
		p.Send("I am an outbound peer\n")
	}
}

func NewServer(addr string) *Server {
	return &Server{addr: addr}
}
