package api

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "goapi/generated/hello"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.SayHelloReq) (*pb.SayHelloRes, error) {
	log.Printf("Received: %v", in.GetSpeech())
	return &pb.SayHelloRes{Speech: "Hello " + in.GetSpeech()}, nil
}

func ServerGrpc() {
	listen, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Grpc Server started on port 50052")

	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
