package main

import (
	pb "cristianrb/calculator/proto"
	"log"
)

func (s *Server) Primes(in *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes times function was invoked with %v\n", in)

	calculatePrimes(int(in.Number), stream)

	return nil
}

func calculatePrimes(n int, stream pb.CalculatorService_PrimesServer) {
	k := 2
	for n > 1 {
		if n%k == 0 {
			stream.Send(&pb.PrimesResponse{Number: uint64(k)})
			n = n / k
		} else {
			k = k + 1
		}
	}
}
