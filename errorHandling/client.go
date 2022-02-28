package main

import (
	"context"
	"fmt"
	"github.com/darkCavalier11/grpc_go/errorHandling/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
)

func main() {
	c, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Unable to dial to the port %v", err)
	}
	findSqrt(c, 41)
	findSqrt(c, -9)

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
