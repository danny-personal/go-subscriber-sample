package main

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/pubsub/v1"
)

func main() {
	subscriptionName := "projects/junior-engineers-gym-2023/subscriptions/urakawa-pubsub-sub"
	keyJson := os.Getenv("KEY_JSON")
	if keyJson == "" {
		fmt.Println("There's no KEY_JSON.")
		return
	} else {
		fmt.Println("KEY_JSON: OK")
	}
	ctx := context.Background()
	client, err := pubsub.NewService(ctx, option.WithCredentialsJSON([]byte(keyJson)))
	if err != nil {
		fmt.Println("Failed to create Client.")
	}
	req := &pubsub.PullRequest{
		MaxMessages: 10,
	}
	res, err := client.Projects.Subscriptions.Pull(subscriptionName, req).Do()
	if err != nil {
		fmt.Println("Failed to pull subscription.")
		return
	}

	for _, msg := range res.ReceivedMessages {
		fmt.Printf("message: %q\n", msg.Message.Data)
		ackReq := &pubsub.AcknowledgeRequest{
			AckIds: []string{msg.AckId},
		}
		if _, err := client.Projects.Subscriptions.Acknowledge(subscriptionName, ackReq).Do(); err != nil {
			//fmt.Errorf("Failed to acknowledge: %v", err)
			fmt.Println(err)
		}
	}
}
