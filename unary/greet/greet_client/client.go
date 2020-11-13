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

	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greet{
			FirstName: "Mahmoud",
			LastName:  "Mekki",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}
