package p2p


//Handshake Func
type HandshakeFunc func (any) error

func NOPHandshakeFunc(any) error {
	return nil
}

func OnPeer(p Peer) error {
	p.Close()
	return nil
}