package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/grpc-course/avercalc/avercalcpb"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("I'am a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}

	defer cc.Close()
	c := avercalcpb.NewAvercalcServiceClient(cc)
	requests := []*avercalcpb.AvercalcRequest{
		{
			Number: 1,
		},
		{
			Number: 2,
		},
		{
			Number: 3,
		},
		{
			Number: 4,
		},
	}
	stream, err := c.Avercalc(context.Background())

	if err != nil {
		log.Printf("Error while making stream %V", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(500 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from server %v", err)
	}
	fmt.Printf("Response: %v\n", res)
}
