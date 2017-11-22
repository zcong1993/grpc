package client

import (
	"context"
	"fmt"
	pb "github.com/zcong1993/grpc/echo"
	"io"
	"log"
	"os"
)

const (
	defaultName = "zcong1993"
	age         = 18
)

func Run() {
	c, conn, err := CreateEchoClient("")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.Echo(context.Background(), &pb.EchoRequest{name, age})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("%+v\n", r)
	r, err = c.EchoAgain(context.Background(), &pb.EchoRequest{name, age})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("%+v\n", r)
	stream, err := c.EchoStream(context.Background(), &pb.EchoRequest{name, age})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		fmt.Printf("%+v\n", r)
	}
}
