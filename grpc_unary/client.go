package main

import (
	"context"
	"fmt"
	"github.com/darkCavalier11/grpc_go/grpc_unary/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	cc, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Unable to connect to the server %v", err)
	}
	defer cc.Close()
	c := gen.NewUnaryServiceClient(cc)
	req := &gen.UnaryRequest{
		Request: "Hello Server. This is client",
	}
	res, err := c.Unary(context.Background(), req)
	if err != nil {
		log.Fatalf("Unable to make request to the server. %v", err)
	}
	fmt.Println(res.GetResponse())
}
