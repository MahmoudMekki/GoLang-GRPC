package main

import (
	"context"
	"log"
	"net"

	"../libpb"
	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *libpb.GreetRequest) (*libpb.GreetResponse, error) {
	first := req.GetGreeting().GetFirstName()
	res := &libpb.GreetResponse{
		Result: "Hello " + first + " !",
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln(err)
	}
	creds, err := credentials.NewServerTLSFromFile("../auth/server.crt", "../auth/server.pem")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer(grpc.Creds(creds))

	libpb.RegisterGreetingServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}

}
