package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/grpc-course/maxfind/maxfindpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Maxfind(stream maxfindpb.MaxfindService_MaxfindServer) error {
	var (
		numbers = make([]int32, 0)
		max     int32
	)

	for {
		req, err := stream.Recv()
		fmt.Printf("Received number: %v\n", req.GetNumreq())
		if err == io.EOF {
			fmt.Println("End of stream")
			return nil
		} else if err != nil {
			log.Fatalf("unable to read streem %v", err)
			return err
		} else {
			numbers = append(numbers, req.GetNumreq())
			for _, v := range numbers {
				if max < v {
					max = v
				}

				err = stream.Send(&maxfindpb.MaxfindResponse{
					Numres: max,
				})
				if err != nil {
					log.Fatalf("error while sending stream to client: %v", err)
					return err
				}
			}
		}
	}
}
func main() {
	fmt.Println("Server bidi started...")
	lis, err := net.Listen("tcp", ":50051")
	fmt.Println("listening port 50051")

	if err != nil {
		log.Fatalf("Server cannot listen on port 50051 %v", err)
	}
	s := grpc.NewServer()

	maxfindpb.RegisterMaxfindServiceServer(s, &server{})

	defer lis.Close()

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
