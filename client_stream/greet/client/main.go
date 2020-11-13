package main

import (
	"context"
	"fmt"
	"log"

	"../greetpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)
	stream, err := c.Greet(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	req := []*greetpb.GreetingRequest{
		&greetpb.GreetingRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Mahmoud",
				LastName:  "Mekki",
			},
		},
		&greetpb.GreetingRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Ola",
				LastName:  "Mekki",
			},
		},
	}
	for _, v := range req {
		stream.Send(v)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%v\n", res)
}
