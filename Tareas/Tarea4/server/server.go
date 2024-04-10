package main

import (
	"context"
	"database/sql"
	"fmt"
	"net"

	pb "serverG/proto"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

// Database configuration

const (
	DBUser     = "root"
	DBPassword = "strong_password"
	DBName     = "tarea4"
	DBHost     = "localhost"
	DBPort     = "3307"
)

type server struct {
	pb.UnimplementedGreeterServer
	db *sql.DB
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	// Store the message in the database
	_, err := s.db.Exec("INSERT INTO music_liderboard (name, album, year, ranking) VALUES (?, ?, ?, ?)",
		in.Name, in.Album, in.Year, in.Rank)
	if err != nil {
		return nil, err
	}

	return &pb.HelloResponse{Message: "Hello I received the data of " + in.Name}, nil
}

func main() {
	// Connect to MySQL database
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUser, DBPassword, DBHost, DBPort, DBName))
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		return
	}
	defer db.Close()

	// Create a listener on TCP port 3001
	lis, err := net.Listen("tcp", ":3001")
	if err != nil {
		fmt.Printf("Failed to listen: %v\n", err)
		return
	}
	fmt.Println("Server listening on port 3001")

	// Create a gRPC server instance
	s := grpc.NewServer()

	// Register the service implementation with the server
	pb.RegisterGreeterServer(s, &server{db: db})

	// Start the server
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v\n", err)
		return
	}
}
