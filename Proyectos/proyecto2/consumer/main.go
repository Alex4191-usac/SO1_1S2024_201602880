package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

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

	mongoClient = initMongoClient()

	var wg sync.WaitGroup

	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			fmt.Println("Error reading message", err)
			break
		}
		fmt.Println("Message received: ", string(msg.Value))

		wg.Add(2)

		go func(msg kafka.Message) {
			defer wg.Done()
			if err := saveToMongo(msg); err != nil {
				fmt.Println("Error saving to MongoDB: ", err)
			}
		}(msg)

		go func(msg kafka.Message) {
			defer wg.Done()
			saveToRedis(msg)
		}(msg)

	}

}

func initMongoClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://35.202.31.252:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
	}
	return client
}

func saveToMongo(msg kafka.Message) error {
	fmt.Println("Saving to Mongo: ")
	collection := mongoClient.Database("sopes1p2").Collection("votossopes")

	doc := map[string]interface{}{
		"message": string(msg.Value),
	}

	_, err := collection.InsertOne(context.Background(), doc)
	if err != nil {
		return fmt.Errorf("error saving to MongoDB: %w", err)
	}

	fmt.Println("Saved to MongoDB")
	return nil

}

func saveToRedis(msg kafka.Message) {
	fmt.Println("Saving to Redis:  ", string(msg.Value))
}
