package main

import (
	"io"
	"log"
	"net"

	"../greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(stream greetpb.GreetService_GreetServer) error {
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetpb.GreetingResponse{Result: result})
		}
		if err == nil {
			firstname := req.GetGreeting().GetFirstName()
			result += " Hello " + firstname + " !"
		} else {
			break
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
	greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
