package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/grpc-course/divsearch/divsearchpb"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("I'am a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}

	defer cc.Close()
	c := divsearchpb.NewDivsearchServiceClient(cc)

	resStream, err := c.DivSearch(context.Background(), &divsearchpb.DivsearchRequest{Number: 240})

	if err != nil {
		log.Printf("Erros while taking result %V", err)
	}

	for {
		res, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("unable to get result: %v", err)
		}
		log.Println(res.GetResult())
	}
}
