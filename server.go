package main

import (
	"context"
	"log"
	"net"

	pb "github.com/nardiyansah/chat/chat"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedChatServiceServer
}

func (s *server) SendChat(ctx context.Context, req *pb.Chat) (*pb.Chat, error) {
	log.Printf("receive message from client : %v", req.Message)
	return &pb.Chat{
		Message: "Hello from server",
	}, nil
}

func main() {

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("error to listen on port 9000 : %v", err)
	}
	defer listener.Close()

	s := grpc.NewServer()

	pb.RegisterChatServiceServer(s, &server{})

	log.Printf("server listening on port %v", listener.Addr())

	if err := s.Serve(listener); err != nil {
		log.Fatalf("grpc server error listening : %v", err)
	}

}
