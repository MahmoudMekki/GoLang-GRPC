package main

import (
	"io"
	"log"
	"net"

	"../greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greeting(stream greetpb.Greet_GreetingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalln(err)
			return err
		}
		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName + " !"
		res := &greetpb.GreetResponse{
			Result: result,
		}
		err = stream.Send(res)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()

	greetpb.RegisterGreetServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
