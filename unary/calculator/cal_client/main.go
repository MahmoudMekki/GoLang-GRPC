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
	c := calpb.NewSumClient(conn)

	req := &calpb.SumRequest{
		Number1: 5,
		Number2: 6,
	}
	res, err := c.Adder(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}
