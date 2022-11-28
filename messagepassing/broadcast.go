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
	broadcastmsg := stream.Chatmessage{
		Messagecontent: host.ID().String(),
		Messagefrom:    host.ID(),
		Authorname:     host.ID().Pretty()[len(host.ID().Pretty())-6 : len(host.ID().Pretty())],
	}
	boradcastbytes, err := json.Marshal(broadcastmsg)
	if err != nil {
		fmt.Println("Error while marshalling the broadcast message")
	}
	for {
		time.Sleep(5 * time.Second)
		err = topic.Publish(ctx, boradcastbytes)
		if err != nil {
			fmt.Println("Error while publishing the broadcast message")
		}
	}
}
