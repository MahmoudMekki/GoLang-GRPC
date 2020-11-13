package main

import (
	"log"
	"net"
	"strconv"
	"time"

	"../greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(req *greetpb.GreetingRequest, stream greetpb.GreetService_GreetServer) error {
	firstname := req.GetGreeting().GetFirstName()

	for i := 0; i < 10; i++ {
		greeting := "Hello " + firstname + " number " + strconv.Itoa(i)
		res := &greetpb.GreetingResponse{
			Result: greeting,
		}
		stream.Send(res)
		time.Sleep(1 * time.Second)
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
