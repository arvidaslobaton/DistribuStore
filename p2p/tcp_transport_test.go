package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAddr := ":3000"
	opts := TCPTransportOpts{
		ListenAddr: ":3000",
		HandshakeFunc: NOPHandshakeFunc,
		Decoder: DefaultDecoder{},
	}

	tr := NewTCPTransport(opts)

	assert.Equal(t, tr.ListenAddr, listenAddr)

	assert.Nil(t, tr.ListenAndAccept())

}

