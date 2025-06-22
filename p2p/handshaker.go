package p2p

//Handshake Func
type HandshakeFunc func (any) error

func NOPHandshakeFunc(any) error {
	return nil
}