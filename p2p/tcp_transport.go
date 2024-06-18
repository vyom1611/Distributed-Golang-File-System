package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPTransport struct {
	ListenAddress string
	listener      net.Listener
	lock          sync.RWMutex
	peers         map[net.Addr]Peer
}

func NewTCPTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
		ListenAddress: listenAddress,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddress)
	if err != nil {
		return err
	}
	t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Println("TCP listener accept error:", err)
		}
		go t.HandleConnection(conn)
	}
}

func (t *TCPTransport) HandleConnection(conn net.Conn) {
	fmt.Println("New incoming connection from:", conn)
}
