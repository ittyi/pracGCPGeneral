package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
)

type UpdateLabelRelationJobEventMessage struct {
	Event   string `json:"event"`
	Action  string `json:"action"`
	TeamID  string `json:"teamId"`
	LabelID string `json:"labelId"`
}

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println("fail to write file")
	}
	publish(f, "", "my-topic", "Hello World")
	log.Println("done")
}

func publish(w io.Writer, projectID, topicID, msg string) error {
	// projectID := "my-project-id"
	// topicID := "my-topic"
	// msg := "Hello World"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub: NewClient: %v", err)
	}
	defer client.Close()

	defaultTopicMessageTeamID := ""
	defaultTopicMessageLabelID := ""

	var testMsg UpdateLabelRelationJobEventMessage
	testMsg.Event = "label"
	testMsg.Action = "create"
	testMsg.Event = defaultTopicMessageTeamID
	testMsg.Event = defaultTopicMessageLabelID

	message, err := json.Marshal(&testMsg)
	if err != nil {
		log.Printf("json.Marshal: %v\n", err)
	}

	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(message),
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("pubsub: result.Get: %v", err)
	}
	fmt.Fprintf(w, "Published a message; msg ID: %v\n", id)
	return nil
}
