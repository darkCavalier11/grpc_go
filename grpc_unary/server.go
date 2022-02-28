package main

import (
	"context"
	"fmt"
	"github.com/darkCavalier11/grpc_go/grpc_unary/gen"
	"google.golang.org/grpc"
	"log"
	"net"
)

type unaryServer struct{}

func (*unaryServer) Unary(ctx context.Context, req *gen.UnaryRequest) (*gen.UnaryResponse, error) {
	ping := req.GetRequest()
	fmt.Println(ping)
	pong := "Request Received 👋"
	res := &gen.UnaryResponse{
		Response: pong,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error listening tcp server %v", err)
	}
	s := grpc.NewServer()
	gen.RegisterUnaryServiceServer(s, &unaryServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Unable to initiate gRPC server %v", err)
	}
}
