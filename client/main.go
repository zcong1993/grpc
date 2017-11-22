package main

import (
	"google.golang.org/grpc"
	"log"
	pb "github.com/zcong1993/grpc/echo"
	"os"
	"context"
	"fmt"
)

const (
	address     = "localhost:9393"
	defaultName = "zcong"
	age	= 18
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEchoerClient(conn)
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.Echo(context.Background(), &pb.EchoRequest{name, age})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("%+v", r)
	r, err = c.EchoAgain(context.Background(), &pb.EchoRequest{name, age})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("%+v", r)
}
