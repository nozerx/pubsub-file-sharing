package p2pnet

import (
	"context"
	"fmt"

	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	drouting "github.com/libp2p/go-libp2p/p2p/discovery/routing"
	dutil "github.com/libp2p/go-libp2p/p2p/discovery/util"
)

func DiscoverPeers(ctx context.Context, host host.Host, service string, kad_dht *dht.IpfsDHT) {
	routingDiscovery := drouting.NewRoutingDiscovery(kad_dht)
	dutil.Advertise(ctx, routingDiscovery, service)
	fmt.Println("Successfull in advertising service")
	connectedPeers := []peer.AddrInfo{}
	isAlreadyConnected := false
	for len(connectedPeers) < 5 {
		fmt.Println("Currently connected to", len(connectedPeers), "out of 5 [for service", service, "]")
		fmt.Println("TOTAL CONNECTIONS : ", len(host.Network().Conns()))
		peerChannel, err := routingDiscovery.FindPeers(ctx, service)
		if err != nil {
			fmt.Println("Error while finding some peers for service :", service)
		} else {
			fmt.Println("Successfull in finding some peers")
		}
		for peerAddr := range peerChannel {

			if peerAddr.ID == host.ID() {
				continue
			}
			for _, connPeers := range connectedPeers {
				if connPeers.ID == peerAddr.ID {
					fmt.Println("Already have a connection with ", peerAddr.ID)
					isAlreadyConnected = true
					break
				}
			}
			if isAlreadyConnected {
				isAlreadyConnected = false
				continue
			}

			err := host.Connect(ctx, peerAddr)
			if err != nil {
				fmt.Println("Error while connecting to peer ", peerAddr.ID)
			} else {
				fmt.Println("Successfull in connecting to peer :", peerAddr.ID)
				connectedPeers = append(connectedPeers, peerAddr)
			}
		}
	}
}
