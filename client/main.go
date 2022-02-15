package client

import (
	"log"

	chat "github.com/hajsf/grpc/chat"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *chat.MessageRequest) (*chat.MessageResponse, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &chat.MessageResponse{Body: "Hello From the Server!"}, nil
}
