package p2p

import (
	"fmt"
	"net"
	"sync"
)

// PUBLIC FUNCTIONS AT TOP, PRIVATE AT BOTTOM
// Order based on importance

/*
TCPPeer represents the remote node over a TCP established connection
Inbound peer = remote node initiates TCP connection w local node.
(Local node listens for incoming connections on a specified port)

Outbound peer = local node actively initiates TCP connection w remote node.
*/

type TCPPeer struct {
	// conn is the underlying connection of the peer
	conn net.Conn

	// if we dial and retrieve a connection => outbound == true
	// if we accept + retrieve a connection => outbound == false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn,
		outbound,
	}
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener // Waits for incoming network connections on a specified address + port
	handshakeFunc HandshakeFunc
	decoder       Decoder

	mu    sync.RWMutex // Put mutex above thing that want to protect
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
		handshakeFunc: NOPHandshakeFunc, //func(any) error { return nil },
		listenAddress: listenAddress,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	// have to use var here as we are not using := for declaration (as t.listener is already declared, and is instead being assigned)
	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()
	return nil

}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept() // blocks until incoming connection is acccepted by listener
		if err != nil {
			fmt.Printf("TCP transport failed to accept connection: %v\n", err)
		}

		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.shakeHands(conn); err != nil {

	}

	for {

	}

	fmt.Printf("new incoming connection: %+v", peer) // %+v prints the struct with field names
}
