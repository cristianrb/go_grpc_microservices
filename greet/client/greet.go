package main

import (
	"context"
	pb "cristianrb/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("do Greet was invoked")
	req := pb.GreetRequest{
		FirstName: "Cristian",
	}

	res, err := c.Greet(context.Background(), &req)
	if err != nil {
		log.Fatalf("Could not greet %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
