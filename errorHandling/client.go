package main

import (
	"context"
	"fmt"
	"github.com/darkCavalier11/grpc_go/errorHandling/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func main() {
	c, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Unable to dial to the port %v", err)
	}
	findExpensiveSqrt(c)

}

func findSqrt(c *grpc.ClientConn, num float64) {
	cc := gen.NewSimpleCalculatorServiceClient(c)
	req, err := cc.Sqrt(context.Background(), &gen.SqrtRequest{ReqNum: num})
	if err != nil {
		resErr, ok := status.FromError(err)
		// if ok, its an actual error from gRPC
		if ok {
			fmt.Println(resErr.Code())
			fmt.Println(resErr.Message())
			return
		} else {
			log.Fatalf("Unable to send request %v", err)
		}
	}
	res := req.GetResNum()
	fmt.Println(res)
}

func findExpensiveSqrt(c *grpc.ClientConn) {
	cc := gen.NewTimeConsumingServiceClient(c)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	req, err := cc.ExpensiveSqrt(ctx, &gen.SqrtRequest{ReqNum: 44})
	if err != nil {
		resErr, ok := status.FromError(err)
		// if ok, its an actual error from gRPC
		if ok {
			fmt.Println(resErr.Code())
			fmt.Println(resErr.Message())
			return
		} else {
			log.Fatalf("Unable to send request %v", err)
		}
	}
	res := req.GetResNum()
	fmt.Println(res)
}
