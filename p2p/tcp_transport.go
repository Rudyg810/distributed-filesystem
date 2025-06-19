package p2p

import (
	"bytes"
	"fmt"
	"net"
	"sync"
)

//This represent the remote nodes over a TCP Connection
type TCPPeer struct {
	// conn is underlying connection of peer
	conn net.Conn
	//if we dial & retrieve a connection => outbound == true 
	outbound bool
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	shakeHands HandshakeFunc
	decoder Decoder
	peerLock      sync.RWMutex
	peers         map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		shakeHands: NOPHandshakeFunc,
		listenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}
	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Print("✖ TCP accept error → ", err, "\n")
			continue
		}
		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn,true)

	if err:= t.shakeHands(conn); err != nil {
		
	}	
	buf  := new(bytes.Buffer)
	for {
		n, _ := conn.Read(buf)
	}
	fmt.Print("🔗 New incoming connection → ", conn.RemoteAddr(), "\n")
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn,
		outbound,
	}
}

