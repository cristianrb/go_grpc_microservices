package main

import (
	"context"
	pb "cristianrb/calculator/proto"
	"log"
)

func sum(c pb.CalculatorServiceClient) {
	log.Println("sum was invoked")
	req := pb.SumRequest{
		First:  4,
		Second: 5,
	}

	res, err := c.Sum(context.Background(), &req)
	if err != nil {
		log.Fatalf("Could not sum %v\n", err)
	}

	log.Printf("Result: %d\n", res.Result)
}
