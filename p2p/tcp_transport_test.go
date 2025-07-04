package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAddr := ":3001"
	opts := TCPTransportOpts{
		ListenAddr: listenAddr,
	}
	tr := NewTCPTransport(opts)
	assert.Equal(t, tr.ListenAddr,listenAddr)
}

