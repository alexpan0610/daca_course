package p2p

import (
	"fmt"
	"net"
)

// Peer represents a peer in the peer-to-peer network.
type Peer struct {
	conn net.Conn
}

func (p *Peer) read() {
	cache := make([]byte, 1024)
	for {
		n, err := p.conn.Read(cache)
		if err != nil {
			continue
		}

		msg := string(cache[:n])
		fmt.Printf("Receive messge %s", msg)
	}
}

// Send push a message to the peer.
func (p *Peer) Send(msg string) {
	_, err := p.conn.Write([]byte(msg))
	fmt.Println(err)
}

func (p *Peer) Start() {
	go p.read()
}

// NewPeer creates a new Peer instance.
func NewPeer(conn net.Conn) *Peer {
	return &Peer{conn: conn}
}
