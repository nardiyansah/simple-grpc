package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/nardiyansah/chat/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error connect to grpc with port 9000 : %v", err)
	}
	defer conn.Close()

	c := pb.NewChatServiceClient(conn)
	ctx := context.Background()

	respon, err := c.SendChat(ctx, &pb.Chat{Message: "hello from clien"})
	if err != nil {
		log.Fatalf("error when call grpc function : %v", err)
	}

	fmt.Printf(respon.Message)
}
