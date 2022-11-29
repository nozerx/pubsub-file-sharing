package main

import (
	flg "pubsubfilesharing/flags"
	pnet "pubsubfilesharing/p2pnet"
	str "pubsubfilesharing/stream"

	"github.com/libp2p/go-libp2p/core/protocol"
)

const service string = "fshr/p2p/rex/trial"
const topic string = "rex/filegroup/group1"

var pid string = "/pid/file/share"

func main() {
	ctx, host := pnet.EstablishP2P()
	host.SetStreamHandler(protocol.ID(pid), str.HandleInputStream)
	kad_dht := pnet.HandleDHT(ctx, host)
	sub, top := pnet.HandlePubSub(ctx, host, topic)
	go pnet.DiscoverPeers(ctx, host, service, kad_dht)
	flg.ResolveAll(ctx, host, top)
	str.HandlePubSubMessages(ctx, host, sub, top)
}
