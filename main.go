package main

import (
	pnet "pubsubfilesharing/p2pnet"
)

const service string = "fshr/p2p/rezon"
const topic string = "rex/filegroup/group1"

func main() {
	ctx, host := pnet.EstablishP2P()
	kad_dht := pnet.HandleDHT(ctx, host)
	pnet.HandlePubSub(ctx, host, topic)
	go pnet.DiscoverPeers(ctx, host, service, kad_dht)
	select {}
}
