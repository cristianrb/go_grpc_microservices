package main

import (
	"context"
	pb "cristianrb/calculator/proto"
	"io"
	"log"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax was invoked")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.MaxRequest{
		{Number: 1},
		{Number: 5},
		{Number: 3},
		{Number: 6},
		{Number: 2},
		{Number: 20},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			stream.Send(req)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}

			log.Printf("Result: %d\n", res.Result)
		}
	}()

	<-waitc
}
