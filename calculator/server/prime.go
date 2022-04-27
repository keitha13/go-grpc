package main

import (
	pb "go_grpc/calculator/proto"
	"log"
)

func (s *Server) Prime(in *pb.PrimeRequest, stream pb.PrimeService_PrimeServer) error {
	log.Printf("Prime function was invoked with: %v\n", in)

	n := in.Num
	d := int64(2)

	for n > 1 {
		if n%d == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: d,
			})
			n /= d
		} else {
			d++
		}
		//res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)

	}
	return nil
}
