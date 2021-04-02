package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"

	"github.com/grpc-course/maxfind/maxfindpb"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("I'am a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}

	defer cc.Close()
	c := maxfindpb.NewMaxfindServiceClient(cc)
	stream, err := c.Maxfind(context.Background())

	if err != nil {
		log.Printf("Error while making stream %V", err)
	}
	waitc := make(chan struct{})
	go func() {

		for i := 0; i < 10; i++ {

			err := stream.Send(&maxfindpb.MaxfindRequest{
				Numreq: getRandNumFrom(100),
			})
			if err != nil {
				log.Fatalf("Erroe during send stream %v", err)
			}
			time.Sleep(400 * time.Millisecond)
		}

		stream.CloseSend()
	}()
	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatalf("error while receiving response from server %v", err)
				break
			}
			fmt.Printf("Response: %v\n", res.GetNumres())
		}
		close(waitc)
	}()
	<-waitc
}
func getRandNumFrom(x int) int32 {
	time.Sleep(time.Millisecond * 50)
	rand.Seed(time.Now().UnixNano())
	return int32(rand.Intn(x))
}
