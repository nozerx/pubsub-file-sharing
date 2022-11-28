package main

import (
	flg "pubsubfilesharing/flags"
	pnet "pubsubfilesharing/p2pnet"
	str "pubsubfilesharing/stream"
)

const service string = "fshr/p2p/rex/trial"
const topic string = "rex/filegroup/group1"

func main() {
	ctx, host := pnet.EstablishP2P()
	kad_dht := pnet.HandleDHT(ctx, host)
	sub, top := pnet.HandlePubSub(ctx, host, topic)
	go pnet.DiscoverPeers(ctx, host, service, kad_dht)
	flg.ResolveAll(ctx, host, top)
	str.HandlePubSubMessages(ctx, host, sub, top)
}
