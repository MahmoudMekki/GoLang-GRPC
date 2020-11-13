package main

import (
	"context"
	"log"
	"net"

	"../greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	firstname := req.GetGreeting().GetFirstName()
	//lastname := req.GetGreeting().GetLastName()
	result := "hello " + firstname

	res := &greetpb.GreetResponse{
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

	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}

}
