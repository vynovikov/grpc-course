package main

import (
	"context"
	"fmt"
	"log"

	"github.com/grpc-course/sumof/sumofpb"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("I'am a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}

	defer cc.Close()
	c := sumofpb.NewSumofServiceClient(cc)

	res, err := c.Sumof(context.Background(), &sumofpb.SumofRequest{Val1: int32(2),
		Val2: int32(1)})

	if err != nil {
		log.Printf("error occured while taking result %V", err)
	}

	log.Printf("result of sum is: %s", res.GetResult())
}
