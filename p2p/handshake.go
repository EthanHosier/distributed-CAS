package p2p

type HandshakeFunc func(any) error

func NOPHandshakeFunc(any) error { // "no operation" handshake function
	return nil
}
