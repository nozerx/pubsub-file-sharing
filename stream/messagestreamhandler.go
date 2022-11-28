package stream

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
)

type chatmessage struct {
	Messagecontent string
	Messagefrom    peer.ID
	Authorname     string
}

func composeMessage(msg string, host host.Host) *chatmessage {
	return &chatmessage{
		Messagecontent: msg,
		Messagefrom:    host.ID(),
		Authorname:     host.ID().Pretty()[len(host.ID().Pretty())-6 : len(host.ID().Pretty())],
	}
}

func readFromSubscription(ctx context.Context, sub *pubsub.Subscription) {
	chatmsg := &chatmessage{}
	for {
		messg, err := sub.Next(ctx)
		if err != nil {
			fmt.Println("Error while getting message from subscription")
		} else {
			err := json.Unmarshal(messg.Data, chatmsg)
			if err != nil {
				fmt.Println("Error while unmarshalling")
			} else {
				fmt.Println(messg.ReceivedFrom.Pretty()[len(messg.ReceivedFrom.Pretty())-6:len(messg.ReceivedFrom.Pretty())], "[", chatmsg.Authorname, "]", ">", string(chatmsg.Messagecontent))
			}
		}
	}

}

func writeToSubscription(ctx context.Context, host host.Host, pubSubTopic *pubsub.Topic) {
	reader := bufio.NewReader(os.Stdin)
	for {
		messg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while reading from standard input")
		} else {
			chatmsg, err := json.Marshal(*composeMessage(messg, host))
			if err != nil {
				fmt.Println("Error while marshalling")
			} else {
				pubSubTopic.Publish(ctx, chatmsg)
			}
		}
	}
}

func HandlePubSubMessages(ctx context.Context, host host.Host, sub *pubsub.Subscription, top *pubsub.Topic) {
	go readFromSubscription(ctx, sub)
	writeToSubscription(ctx, host, top)
}
