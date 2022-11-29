package messagepassing

import (
	"context"
	"encoding/json"
	"fmt"
	"pubsubfilesharing/stream"

	"time"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
)

func BroadCastMentorDetails(ctx context.Context, host host.Host, topic *pubsub.Topic) {
	broadcastmsg := stream.BroadcastMsg{
		MentorNode: host.ID(),
	}
	boradcastbytes, err := json.Marshal(broadcastmsg)
	if err != nil {
		fmt.Println("Error while marshalling the broadcast message")
	}
	packetMsg := stream.Packet{
		Type:         "brd",
		InnerContent: boradcastbytes,
	}
	packetBytes, err := json.Marshal(packetMsg)
	if err != nil {
		fmt.Println("Error while marshalling the broadcast message")
	}
	for {
		time.Sleep(30 * time.Second)
		err = topic.Publish(ctx, packetBytes)
		if err != nil {
			fmt.Println("Error while publishing the broadcast message")
		}
	}
}
