package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "client/proto"
)

var ctx = context.Background()

type Message struct {
	Name  string
	Album string
	Year  string
	Rank  string
}

func insertMessage(c *fiber.Ctx) error {

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	msg := Message{
		Name:  data["name"],
		Album: data["album"],
		Year:  data["year"],
		Rank:  data["rank"],
	}

	go sendRequest(msg)
	return nil
}

func sendRequest(msg Message) {
	conn, err := grpc.Dial("localhost:3001", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())

	if err != nil {
		fmt.Println("Failed to connect:", err)

	}

	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	response, err := client.SayHello(ctx, &pb.HelloRequest{
		Name:  msg.Name,
		Album: msg.Album,
		Year:  msg.Year,
		Rank:  msg.Rank,
	})

	if err != nil {
		fmt.Println("Failed to call SayHello:", err)

	}

	fmt.Println("Server response:", response)

}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "grcp test!"})
	})

	app.Post("/insert", insertMessage)

	app.Listen(":3000")

}
