package main

import (
	"log"

	"github.com/rudyg810/distributedFS/p2p"
)


func main() {
	listenAddr := ":3001"
	opts := p2p.TCPTransportOpts{
		ListenAddr: listenAddr,
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder: nil,
	}
	tr := p2p.NewTCPTransport(opts)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}