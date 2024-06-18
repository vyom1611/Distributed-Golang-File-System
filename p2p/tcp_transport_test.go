package p2p

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	listenerAddr := ":8080"
	transport := NewTCPTransport(listenerAddr)

	assert.Equal(t, transport.ListenAddress, listenerAddr)

	//Server
	assert.Nil(t, transport.ListenAndAccept())

}
