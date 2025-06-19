package main

import (
	"fmt"
	"log"

	"github.com/rudyg810/distributedFS/p2p"
)


func main(){
	port  := ":3001"
	tr := p2p.NewTCPTransport(port)
	if err := tr.ListenAndAccept(); err != nil  {
		log.Fatal(err)
	}
	fmt.Print("TCP Listener started at",port)
	select{}
}