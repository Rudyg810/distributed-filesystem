package main

import (
	"fmt"
	"log"  
	"github.com/rudyg810/distributedFS/p2p"
)


func main() {
	listenAddr := ":3001"
	opts := p2p.TCPTransportOpts{
		ListenAddr: listenAddr,
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder: p2p.DefaultDecoder{},
		OnPeer: p2p.OnPeer,
		
	}
	tr := p2p.NewTCPTransport(opts)
	go func ()  {
		for {
			msg := <-tr.Consume()
			fmt.Println("New message in Channel",len(msg.Payload))
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}