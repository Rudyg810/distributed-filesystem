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
	TCPTransportOpts
	Listener      net.Listener

	peerLock sync.RWMutex
	peer map[net.Addr]Peer
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer  {
	return &TCPPeer{
		conn,
		outbound,
	}
}

type TCPTransportOpts  struct {
	ListenAddr string
	HandshakeFunc HandshakeFunc
	Decoder Decoder
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts:opts,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	ln, err := net.Listen("tcp",t.ListenAddr)
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
	peer := NewTCPPeer(conn,true)
	if err := t.HandshakeFunc(peer); err != nil {
		fmt.Print("TCP Handshake error",err)
		conn.Close()
		return
	}

	//read Loop
	msg := &Temp{}
	for {
		if err := t.Decoder.Decode(conn, msg); err != nil {
			fmt.Print("TCP Error",err)
			continue
		}
	}
}