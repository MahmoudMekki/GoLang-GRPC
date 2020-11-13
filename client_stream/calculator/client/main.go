package main

import (
	"context"
	"fmt"
	"log"

	"../calpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	c := calpb.NewAverageClient(conn)
	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	req := []*calpb.AVGRequest{
		&calpb.AVGRequest{
			Number: 1,
		},
		&calpb.AVGRequest{
			Number: 2,
		},
		&calpb.AVGRequest{
			Number: 3,
		},
		&calpb.AVGRequest{
			Number: 4,
		},
		&calpb.AVGRequest{
			Number: 5,
		},
		&calpb.AVGRequest{
			Number: 6,
		},
		&calpb.AVGRequest{
			Number: 8,
		},
	}
	for _, req := range req {
		stream.Send(req)
	}
	r, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%v\n", r)
}
