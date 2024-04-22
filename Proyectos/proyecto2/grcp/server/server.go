package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grcp_server/proto"

	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
)

type serverImp struct {
	pb.UnimplementedVoteServiceServer
}

type VoteData struct {
	name  string
	album string
	year  string
	rank  string
}

func kafkaVote(vote VoteData) {
	writer := &kafka.Writer{
		Addr:         kafka.TCP("my-cluster-kafka-bootstrap:9092"),
		Topic:        "my-topic",
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
	}

	defer writer.Close()

	//Print the vote data
	message := fmt.Sprintf("KafkaName: %s\nAlbum: %s\nYear: %s\nRank: %s\n", vote.name, vote.album, vote.year, vote.rank)

	// Write the vote data to the kafka topic
	msg := kafka.Message{
		Key:   []byte(vote.name),
		Value: []byte(message),
	}

	if err := writer.WriteMessages(context.Background(), msg); err != nil {
		log.Fatal("Failed to write messages:", err)
	}
	fmt.Println("Vote has been sent to Kafka")
}

func (s *serverImp) TakeVote(ctx context.Context, in *pb.Vote) (*pb.VoteResponse, error) {
	log.Printf("Received: %v", in.GetAlbum())

	// Save the vote data in a struct
	voteData := VoteData{
		name:  in.GetName(),
		album: in.GetAlbum(),
		year:  in.GetYear(),
		rank:  in.GetRank(),
	}

	// Print the vote data
	fmt.Printf("Name: %s\nAlbum: %s\nYear: %s\nRank: %s\n", voteData.name, voteData.album, voteData.year, voteData.rank)

	// Send the vote data to Kafka
	kafkaVote(voteData)
	// Return the response
	return &pb.VoteResponse{Message: "Vote has been received"}, nil
}

func main() {
	fmt.Println("Server Implementation")

	lis, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterVoteServiceServer(s, &serverImp{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
