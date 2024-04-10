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

type server struct {
	pb.UnimplementedGreeterServer
}

type Message struct {
	Name  string
	Album string
	Year  string
	Rank  string
}

type DBconnection struct {
	db *sql.DB
}

func (dbc *DBconnection) Connect(user, password, host, port, database string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	dbc.db = db
	fmt.Println("Connected to the database")
	return nil
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
	fmt.Println("Server listening")

	dbc := DBconnection{}
	err_db := dbc.Connect("root", "strong_password", "localhost", "3307", "tarea4")
	if err_db != nil {
		fmt.Println("Failed to connect to the database:", err_db)
	}

	defer dbc.db.Close()

	if err := s.Serve(lis); err != nil {
		fmt.Println("Failed to serve:", err)
	}

}
