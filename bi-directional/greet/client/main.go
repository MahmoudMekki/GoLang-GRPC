package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"../greetpb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	c := greetpb.NewGreetClient(conn)

	stream, err := c.Greeting(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	reqs := []*greetpb.GreetRequest{
		&greetpb.GreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Mahmoud",
				LastName:  "Mekki",
			},
		},
		&greetpb.GreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Ola",
				LastName:  "Mekki",
			},
		},
	}

	wc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			err := stream.Send(req)
			if err != nil {
				log.Fatalln(err)
			}
			time.Sleep(time.Second * 1)
		}
		stream.CloseSend()

	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				close(wc)
			}
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(res)
		}

	}()

	<-wc
}
