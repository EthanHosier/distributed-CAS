package main

import (
	"log"

	"github.com/ethanhosier/foreverstore/p2p"
)

/*
Send file
Hash it
Add interface to transform hashed file name
Make subfolders (pairs of two) - store the actual encrypted data there
*/

/*
Peer-to-peer (P2P) refers to a decentralized network architecture where participants (peers) interact directly with each other, rather than through a central server or authority. In a P2P network, each participant acts both as a client and a server, sharing resources, services, or data directly with other participants.
*/

func main() {
	tr := p2p.NewTCPTransport(":3000")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}

}
