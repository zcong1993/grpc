package client

import (
	pb "github.com/zcong1993/grpc/echo"
	"google.golang.org/grpc"
)

const addressDefault = "localhost:9393"

func CreateEchoClient(address string) (pb.EchoerClient, *grpc.ClientConn, error) {
	if address == "" {
		address = addressDefault
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}
	c := pb.NewEchoerClient(conn)
	return c, conn, nil
}
