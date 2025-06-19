package p2p

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T)  {
	listenAdd := ":4000"
	tr := NewTCPTransport(listenAdd)
	assert.Equal(t, tr.listenAddress, listenAdd)

	//Server
	assert.Nil(t,tr.ListenAndAccept())
	select {}
}