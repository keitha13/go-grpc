package main

import (
	"context"
	pb "go_grpc/greet/proto"
	"io"
	"log"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")
	req := &pb.GreetRequest{
		FirstName: "Hayashi",
	}
	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed while calling GreetManyTimes: %v\n", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}
		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}
}
