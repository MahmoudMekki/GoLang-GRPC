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

	req := &greetpb.GreetingRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Mahmoud",
			LastName:  "Mekki",
		},
	}
	stream, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		res, err := stream.Recv()
		if err == nil {
			fmt.Println(res)
		} else {
			break
		}

	}

	fmt.Println("Streaming has ended!")

}
