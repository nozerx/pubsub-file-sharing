package flags

import (
	"context"
	"flag"
	msgpass "pubsubfilesharing/messagepassing"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
)

func ResolveAll(ctx context.Context, host host.Host, top *pubsub.Topic) {
	mode := flag.Int("mode", 0, " 0 - for normal mode | 1- for mentor mode")
	flag.Parse()

	if *mode == 1 {
		go msgpass.BroadCastMentorDetails(ctx, host, top)
	}

}
