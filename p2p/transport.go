package p2p

// THis is a peer node in network
type Peer struct {


}
//This is an communication between nodes in network
//TCP UDP Websockets
type Transport interface {
	ListenAndAccept() error
}