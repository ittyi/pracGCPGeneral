package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"sync/atomic"
	"time"

	"cloud.google.com/go/pubsub"
)

func main() {
	f, err := os.Create("./testSubsc.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println("fail to write file")
	}
	pullMsgs(f, "", "my-sub")
	log.Println("done")
}

func pullMsgs(w io.Writer, projectID, subID string) error {
	// projectID := "my-project-id"
	// subID := "my-sub"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	sub := client.Subscription(subID)

	// Receive messages for 10 seconds, which simplifies testing.
	// Comment this out in production, since `Receive` should
	// be used as a long running operation.
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var received int32
	err = sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		fmt.Fprintf(w, "Got message: %q\n", string(msg.Data))
		atomic.AddInt32(&received, 1)
		msg.Ack()
	})
	if err != nil {
		return fmt.Errorf("sub.Receive: %v", err)
	}
	fmt.Fprintf(w, "Received %d messages\n", received)

	return nil
}
