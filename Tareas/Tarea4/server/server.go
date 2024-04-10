package main

import (
	"context"
	"fmt"
	"net"

	pb "serverG/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

type Message struct {
	Name  string
	Album string
	Year  string
	Rank  string
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	data := Message{
		Name:  in.Name,
		Album: in.Album,
		Year:  in.Year,
		Rank:  in.Rank,
	}
	fmt.Println(data)
	return &pb.HelloResponse{Message: "Hello I received the data of " + data.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":3001")
	if err != nil {
		fmt.Println("Failed to listen:", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	fmt.Println("Server listening on port 3001")
	if err := s.Serve(lis); err != nil {
		fmt.Println("Failed to serve:", err)
	}

}
