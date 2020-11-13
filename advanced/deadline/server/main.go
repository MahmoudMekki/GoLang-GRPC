package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"../libpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greeting(ctx context.Context, req *libpb.GreetingRequest) (*libpb.GreetingResponse, error) {

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			fmt.Println(" Client cancelled the request!!!")
			return nil, status.Error(codes.Canceled, "Deadline exceeded!")
		}
		time.Sleep(1 * time.Second)
	}
	first := req.GetGreeting().GetFirstname()
	result := "Hello " + first + " !"
	res := &libpb.GreetingResponse{
		Result: result,
	}
	return res, nil

}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()

	libpb.RegisterGreetServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
