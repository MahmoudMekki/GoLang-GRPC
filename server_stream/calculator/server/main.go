package main

import (
	"log"
	"net"
	"time"

	"../calpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Primes(req *calpb.PrimeRequest, stream calpb.Prime_PrimesServer) error {
	number := req.GetNumber()
	k := 2
	for number > 1 {
		if int(number)%k == 0 {
			res := &calpb.PrimeResponse{
				Prime: int32(k),
			}
			stream.Send(res)
			number /= int32(k)
			time.Sleep(time.Second * 1)
		} else {
			k++
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	calpb.RegisterPrimeServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
