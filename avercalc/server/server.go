package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/grpc-course/avercalc/avercalcpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Avercalc(stream avercalcpb.AvercalcService_AvercalcServer) error {
	var (
		num int
		sum float64
	)

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&avercalcpb.AvercalcResponse{
				Result: sum / float64(num),
			})
		} else if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		} else {
			num++
			sum += float64(req.GetNumber())
		}
	}
}

func main() {

	fmt.Println("Server started...")
	lis, err := net.Listen("tcp", ":50051")
	fmt.Println("listening port 50051")

	if err != nil {
		log.Fatalf("Server cannot listen on port 50051 %v", err)
	}
	s := grpc.NewServer()

	avercalcpb.RegisterAvercalcServiceServer(s, &server{})

	defer lis.Close()

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
