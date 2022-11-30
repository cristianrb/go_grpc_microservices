package main

import (
	pb "cristianrb/calculator/proto"
	"io"
	"log"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg function was invoked")

	res := 0.0
	count := 0.0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{Result: res / count})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		res += float64(req.Number)
		count++
	}
}
