package main

import (
	"log"

	"github.com/arvidaslobaton/DistribuStore/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr: ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder: p2p.GOBDecoder{},

	}

	tr := p2p.NewTCPTransport(tcpOpts)

	err := tr.ListenAndAccept(); if err != nil {
		log.Fatal(err)
	}

	select{}

}