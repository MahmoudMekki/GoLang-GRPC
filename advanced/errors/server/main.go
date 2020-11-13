package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	libpb "../libpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) SquareRoot(ctx context.Context, req *libpb.SquareRootRequest) (*libpb.SquareRootResponse, error) {
	number := req.GetNumber()
	if number <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Recevied invalid number!= %v", number))
	}
	root := math.Sqrt(number)
	res := &libpb.SquareRootResponse{
		Result: root,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	libpb.RegisterSquareRootServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
