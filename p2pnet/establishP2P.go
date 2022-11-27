package p2pnet

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
)

func EstablishP2P() (context.Context, host.Host) {
	host, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"))
	if err != nil {
		fmt.Println("Error while creating a new node")
	} else {
		fmt.Println("Successfully created a new node")
	}
	ctx := context.Background()
	return ctx, host
}