package p2p

//Peer is an interface that represents the remote node
type Peer interface {}

//Transport is anything that that handles the communication
//between nodes in network. This can be of form of (TCP, UDP, WebSockets)
type Transport interface {
	ListenAndAccept() error
}