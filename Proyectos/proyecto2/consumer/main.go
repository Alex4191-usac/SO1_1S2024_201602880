package main

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var redisClient *redis.Client

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

	redisClient = initRedis()
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
			if err := saveToRedis(msg); err != nil {
				fmt.Println("Error saving to Redis: ", err)
			}
		}(msg)

		go func(msg kafka.Message) {
			defer wg.Done()
			if err := saveToMongo(msg); err != nil {
				fmt.Println("Error saving to MongoDB: ", err)
			}
		}(msg)

	}

}

func initRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "10.55.186.172:6379", //34.27.178.49
		Password: "admin",
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
	}

	fmt.Println("Connected to Redis")
	return client
}

func initMongoClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://10.55.185.189:27017") //34.28.96.178
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
	}
	fmt.Println("Connected to Mongo")
	return client
}

func saveToMongo(msg kafka.Message) error {
	fmt.Println("PROCESS TO SAVE to Mongo: ")
	collection := mongoClient.Database("sopes1p2").Collection("votos")

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

func saveToRedis(msg kafka.Message) error {
	fmt.Println("PROCESS TO SAVE to REDIS: ")

	//split message by /n
	splitted := strings.Split(string(msg.Value), "\n")
	fmt.Println("SPLITEO: ")
	name := strings.TrimPrefix(splitted[0], "Name: ")
	album := strings.TrimPrefix(splitted[1], "Album: ")
	year := strings.TrimPrefix(splitted[2], "Year: ")
	rank := strings.TrimPrefix(splitted[3], "Rank: ")

	fmt.Printf("Valores extraídos: name='%s', album='%s', year='%s', rank='%s'\n", name, album, year, rank)

	key := fmt.Sprintf("%s_%s_%s_%s", name, album, year, rank)

	fmt.Printf("Key Generated for reddis: %s\n", key)

	_, err := redisClient.Incr(context.Background(), key).Result()
	if err != nil {
		return fmt.Errorf("error increment the counter on Redis: %w", err)
	}

	fmt.Printf("Saved to Redis %s/n", key)

	return nil

}
