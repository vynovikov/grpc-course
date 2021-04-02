package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/grpc-course/divsearch/divsearchpb"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) DivSearch(req *divsearchpb.DivsearchRequest, stream divsearchpb.DivsearchService_DivSearchServer) error {
	res := divcalc(req.GetNumber())
	for _, v := range res {

		stream.Send(&divsearchpb.DivsearchResponse{
			Result: v,
		})
		time.Sleep(600 * time.Millisecond)
	}

	return nil
}

func divcalc(n int32) []int32 {

	var kk int32 = 2
	k := make([]int32, 0)

	for n > 1 {
		if n%kk == 0 {
			k = append(k, kk)
			n = n / kk
		} else {
			kk++
		}
	}

	return k
}
func main() {

	fmt.Println("Server started...")
	lis, err := net.Listen("tcp", ":50051")
	fmt.Println("listening port 50051")

	if err != nil {
		log.Fatalf("Server cannot listen on port 50051 %v", err)
	}
	s := grpc.NewServer()

	divsearchpb.RegisterDivsearchServiceServer(s, &server{})

	defer lis.Close()

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
