package main

import (
	"context"
	"fmt"
	"log"

	"../libpb"
	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("../auth/ca.crt", "")
	if err != nil {
		log.Fatalln(err)
	}
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalln(err)
	}
	c := libpb.NewGreetingClient(conn)
	req := &libpb.GreetRequest{
		Greeting: &libpb.Greet{
			FirstName: "Mahmoud",
			LastName:  "Mekki",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res.Result)
}
