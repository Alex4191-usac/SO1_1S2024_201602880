package main

import (
	"context"
	"fmt"
	"log"

	pb "grcp_client/proto"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ctx = context.Background()

type VoteData struct {
	name  string
	album string
	year  string
	rank  string
}

//prepare the Messate to send to the server

func prepareVote(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	voteData := VoteData{
		name:  data["name"],
		album: data["album"],
		year:  data["year"],
		rank:  data["rank"],
	}

	go sendVote(voteData)
	return nil

}

func sendVote(voteData VoteData) {
	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	cl := pb.NewVoteServiceClient(conn)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Error closing connection: %v", err)
		}
	}(conn)

	vto, err := cl.TakeVote(ctx, &pb.Vote{
		Name:  voteData.name,
		Album: voteData.album,
		Year:  voteData.year,
		Rank:  voteData.rank,
	})

	if err != nil {
		log.Fatalf("could not send vote: %v", err)
	}

	fmt.Println("Response from the server: ", vto.GetMessage())

}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/vote", prepareVote)

	app.Listen(":3000")
}
