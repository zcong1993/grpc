package main

import (
	"context"
	pb "github.com/zcong1993/grpc/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const port = ":9393"

type server struct{}

func (s *server) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Println("invoke echo")
	return &pb.EchoResponse{Name: in.Name, Age: in.Age}, nil
}

func (s *server) EchoAgain(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Println("invoke echo again")
	return &pb.EchoResponse{Name: in.Name, Age: in.Age}, nil
}

func (s *server) EchoStream(req *pb.EchoRequest, stream pb.Echoer_EchoStreamServer) error {
	for i := 0; i < 10; i++ {
		if err := stream.Send(&pb.EchoResponse{Name: req.Name, Age: req.Age}); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEchoerServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
