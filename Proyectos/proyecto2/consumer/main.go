package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"my-cluster-kafka-bootstrap:9092"},
		Topic:     "my-topic",
		Partition: 0,
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})

	defer reader.Close()

	ctx := context.Background()
	fmt.Println("Starting Kafka consumer")

	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			fmt.Println("Error reading message", err)
			break
		}
		fmt.Println("Message received: ", string(msg.Value))
	}

}
