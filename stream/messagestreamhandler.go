package stream

import (
	"bufio"
	"context"
	"fmt"
	"os"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

func readFromSubscription(ctx context.Context, sub *pubsub.Subscription) {
	for {
		messg, err := sub.Next(ctx)
		if err != nil {
			fmt.Println("Error while getting message from subscription")
		} else {
			fmt.Println(messg.ReceivedFrom.Pretty()[len(messg.ReceivedFrom.Pretty())-6:len(messg.ReceivedFrom.Pretty())], ">", string(messg.Data))
		}
	}

}

func writeToSubscription(ctx context.Context, pubSubTopic *pubsub.Topic) {
	reader := bufio.NewReader(os.Stdin)
	for {
		messg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while reading from standard input")
		} else {
			pubSubTopic.Publish(ctx, []byte(messg))
		}
	}
}

func HandlePubSubMessages(ctx context.Context, sub *pubsub.Subscription, top *pubsub.Topic) {
	go readFromSubscription(ctx, sub)
	writeToSubscription(ctx, top)
}
