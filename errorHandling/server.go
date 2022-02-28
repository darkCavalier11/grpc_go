package main

import (
	"context"
	"fmt"
	"github.com/darkCavalier11/grpc_go/errorHandling/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math"
	"net"
)

type server struct {
}

func (*server) Sqrt(ctx context.Context, req *gen.SqrtRequest) (*gen.SqrtResponse, error) {
	reqNum := req.GetReqNum()
	if reqNum < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Received -ve number")
	}
	res := math.Sqrt(reqNum)
	return &gen.SqrtResponse{ResNum: res}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Unable to listen on address %v", err)
	}
	s := grpc.NewServer()
	gen.RegisterSimpleCalculatorServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Unable to listen gRPC server %v", err)
	}
}
