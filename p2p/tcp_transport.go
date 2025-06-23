package p2p

import (
	"fmt"
	"io"
	"net"

)

type TCPPeer struct {
	conn net.Conn
	outbound bool
}

type TCPTransport struct {
	TCPTransportOpts
	Listener net.Listener
	rpcch chan RPC

	// peerLock sync.RWMutex
	// peer     map[net.Addr]Peer
	OnPeer func(Peer) error
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn,
		outbound,
	}
}

//Auto implements close from Peer
func (p *TCPPeer) Close() error {
	return p.conn.Close()
}

type TCPTransportOpts struct {
	ListenAddr    string
	HandshakeFunc HandshakeFunc
	Decoder       Decoder
	OnPeer func(Peer) error
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		rpcch: make(chan RPC),
	}
}

//Consume implements the transport interface, which will return read only channel
func (t *TCPTransport) Consume() <-chan RPC  {
	return t.rpcch
}

func (t *TCPTransport) ListenAndAccept() error {
	ln, err := net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return err
	}
	t.Listener = ln

	fmt.Println("✅ TCP Server started on", t.ListenAddr) 

	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.Listener.Accept()
		if err != nil {
			fmt.Print("Accept Loop error", err)
		}

		fmt.Println("🔌 New peer connected:", conn.RemoteAddr()) 

		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	var err error
	defer func ()  {
		fmt.Println("Dropping Peer Connection", err)
		conn.Close()
	}()
	peer := NewTCPPeer(conn, true)
	if err := t.HandshakeFunc(peer); err != nil {
		fmt.Print("TCP Handshake error", err)
		conn.Close()
		return
	}

	if t.OnPeer != nil {
		if err := t.OnPeer(peer); err != nil {
			fmt.Println("On Peer Error Dropping")
			return
		}
	}

	rpc := &RPC{}
	for {
		if err := t.Decoder.Decode(conn, rpc); err != nil {
			fmt.Print("TCP Message Error", err)
			if err == io.EOF {
				fmt.Println("🚪 Connection closed by peer:", conn.RemoteAddr())
				break
			}
			continue
		}
		rpc.From = conn.RemoteAddr()
		t.rpcch <-*rpc
		fmt.Println("Message", string(rpc.Payload))
	}
}
