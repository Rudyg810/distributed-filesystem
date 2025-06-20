package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPPeer struct {
	conn net.Conn
	//outbound -> You connect to them
	//inbound -> They connect to you
	// if we dial and retrieve the connection => outbound true
	// if we accept and retrieve the connection => inbound true
	outbound bool
}

type TCPTransport struct {
	listenAddress string
	Listener      net.Listener
	shakeHands HandshakeFunc
	decoder Decoder

	peerLock sync.RWMutex
	peer map[net.Addr]Peer
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer  {
	return &TCPPeer{
		conn,
		outbound,
	}
}

func NewTCPTransport(address string) *TCPTransport {
	return &TCPTransport{
		shakeHands: NOPHandshakeFunc,
		listenAddress: address,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	ln, err := net.Listen("tcp",t.listenAddress)
	if err != nil {
		return err
	}
	t.Listener = ln

	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.Listener.Accept()
		if err != nil {	
			fmt.Print("Error", err)
		}

		go t.handleConn(conn)
	}
}

type Temp struct{}

func (t *TCPTransport) handleConn(conn net.Conn){
	defer conn.Close()
	// peer := NewTCPPeer(conn,false)
	if err := t.shakeHands(conn); err != nil {
		fmt.Print(err)
	}

	//read Loop
	msg := &Temp{}
	for {
		if err := t.decoder.Decode(conn, msg); err != nil {
			fmt.Print("TCP Error",err)
			continue
		}
	}
}