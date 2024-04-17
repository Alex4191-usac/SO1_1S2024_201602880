package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grcp_server/proto"

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
