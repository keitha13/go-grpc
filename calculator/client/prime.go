package main

import (
	"context"
	pb "go_grpc/calculator/proto"
	"io"
	"log"
)

func doPrime(c pb.PrimeServiceClient) {
	log.Println("doPrime was invoked")
	req := &pb.PrimeRequest{
		Num: 120,
	}
	stream, err := c.Prime(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed while calling Prime: %v\n", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}
		log.Printf("Primes: %v\n", msg.Result)
	}
}
