package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/grpc-course/sumof/sumofpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sumof(c context.Context, req *sumofpb.SumofRequest) (*sumofpb.SumofResponse, error) {
	return &sumofpb.SumofResponse{
		Result: strconv.Itoa(int(req.GetVal1() + req.GetVal2())),
	}, nil
}

func main() {

	fmt.Println("Server started...")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Server cannot listen on port 50051 %v", err)
	}
	s := grpc.NewServer()

	sumofpb.RegisterSumofServiceServer(s, &server{})

	defer lis.Close()

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
