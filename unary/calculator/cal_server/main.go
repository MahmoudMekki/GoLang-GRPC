package main

import (
	"context"
	"log"
	"net"

	"../calpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Adder(ctx context.Context, req *calpb.SumRequest) (*calpb.SumResponse, error) {
	n1 := req.GetNumber1()
	n2 := req.GetNumber2()
	sum := n1 + n2
	res := &calpb.SumResponse{
		Result: sum,
	}
	return res, nil

}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	calpb.RegisterSumServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}

}
