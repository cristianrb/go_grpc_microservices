package main

import (
	pb "cristianrb/calculator/proto"
	"io"
	"log"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max was invoked")
	max := uint64(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		if req.Number > max {
			max = req.Number
			err = stream.Send(&pb.MaxResponse{
				Result: max,
			})
			if err != nil {
				log.Fatalf("Error while sending data to client: %v\n", err)
			}
		}
	}
}
