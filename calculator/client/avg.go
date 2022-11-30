package main

import (
	"context"
	pb "cristianrb/calculator/proto"
	"log"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg was invoked")

	reqs := []*pb.AvgRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}

	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Avg: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from Avg: %v\n", err)
	}

	log.Printf("Avg: %f\n", res.Result)
}