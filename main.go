package main

import (
	"fmt"
	"github.com/elastos/cource/p2p"
)

func main() {
	svr := p2p.NewServer("localhost:30000")
	go svr.Start()

	addrs := make([]string, 0, 10)
	for i := 0; i < 10; i++ {
		addr := fmt.Sprint("localhost:30000")
		addrs = append(addrs, addr)
	}

	svr.Connect(addrs)

	select {}
}
