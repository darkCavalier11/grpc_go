package main

import (
	"context"
	"fmt"
	"github.com/darkCavalier11/grpc_go/grpc_streaming/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"time"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Unable to connect to the server %v", err)
	}
	defer cc.Close()
	bidirectionalStreamHandler(cc)
}

func bidirectionalStreamHandler(cc *grpc.ClientConn) {
	c := gen.NewBiDirectionalStreamingServiceClient(cc)
	stream, err := c.BiDirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Unable to make request to the server. %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for i := 0; i < 10; i++ {
			req := &gen.SimpleRequest{
				Request: fmt.Sprintf("Hello from client, count %v \U0001FA83", i),
			}
			err := stream.Send(req)
			if err != nil {
				log.Fatalf("Unable to send stream to server %v", err)
			}
			time.Sleep(time.Second)
		}
		if err := stream.CloseSend(); err != nil {
			fmt.Printf("Unable to close client stream %v", err)
		}
	}()
	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				break
			}
			if err != nil {
				log.Fatalf("Error occured during server streaming %v", err)
			}
			fmt.Println(msg.GetResponse())
		}
	}()
	<-waitc
}

func clientStreamHandler(cc *grpc.ClientConn) {
	c := gen.NewClientStreamingServiceClient(cc)
	stream, err := c.ClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Unable to make request to the server. %v", err)
	}
	for i := 0; i < 10; i++ {
		req := &gen.SimpleRequest{
			Request: fmt.Sprintf("Hello from client, count %v \U0001FA83", i),
		}
		err := stream.Send(req)
		if err != nil {
			log.Fatalf("Unable to send stream to server %v", err)
		}
		time.Sleep(time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Printf("Unable receive response %v", err)
	}
	fmt.Println(res.GetResponse())
}

func serverStreamingHandler(cc *grpc.ClientConn) {
	c := gen.NewServerStreamingServiceClient(cc)
	req := &gen.SimpleRequest{
		Request: "Hello Server. This is client",
	}
	res, err := c.ServerStreaming(context.Background(), req)
	if err != nil {
		log.Fatalf("Unable to make request to the server. %v", err)
	}
	for {
		msg, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error occured during server streaming %v", err)
		}
		fmt.Println(msg.GetResponse())
	}
}
