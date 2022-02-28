package main

import (
	"fmt"
	"github.com/darkCavalier11/grpc_go/grpc_streaming/gen"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"time"
)

type server struct{}

func (*server) ServerStreaming(req *gen.SimpleRequest, stream gen.ServerStreamingService_ServerStreamingServer) error {
	r := req.GetRequest()
	fmt.Println(r)
	for i := 0; i < 10; i++ {
		err := stream.Send(&gen.SimpleResponse{
			Response: fmt.Sprintf("Hello. This is no %v 👋🏻 from server", i),
		})
		if err != nil {
			log.Fatalf("Unable to streaming %v", err)
		}
		time.Sleep(time.Second)
	}
	return nil
}

func (*server) ClientStreaming(stream gen.ClientStreamingService_ClientStreamingServer) error {
	result := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Unable to receive stream from client %v", err)
		}
		fmt.Println(req.GetRequest())
		result += 1
	}
	return stream.SendAndClose(&gen.SimpleResponse{
		Response: fmt.Sprintf("Received %v requests", result),
	})
}

func (*server) BiDirectionalStreaming(stream gen.BiDirectionalStreamingService_BiDirectionalStreamingServer) error {
	result := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading the stream from client %v", err)
		}
		fmt.Println(req.GetRequest())
		err = stream.Send(&gen.SimpleResponse{
			Response: fmt.Sprintf("Received request no %v", result),
		})
		result += 1
		if err != nil {
			log.Fatalf("Unable to send stream to client %v", err)
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Unable start server %v", err)
	}
	s := grpc.NewServer()
	gen.RegisterServerStreamingServiceServer(s, &server{})
	gen.RegisterClientStreamingServiceServer(s, &server{})
	gen.RegisterBiDirectionalStreamingServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Unable to bind with grpc %v", err)
	}
}
