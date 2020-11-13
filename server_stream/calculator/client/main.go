package main

import (
	"context"
	"fmt"
	"io"
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

	c := calpb.NewPrimeClient(conn)
	req := &calpb.PrimeRequest{
		Number: 120,
	}
	p, err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		res, err := p.Recv()
		if err == io.EOF {
			log.Fatalln(err)
		}
		if err == nil {
			fmt.Println(res)
		} else {
			break
		}
	}
	fmt.Println("Streaming has been ended!")
}
