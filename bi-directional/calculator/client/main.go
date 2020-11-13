package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"../calpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	c := calpb.NewMaxClient(conn)
	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	reqs := []*calpb.Maxrequest{
		&calpb.Maxrequest{
			Number: 1,
		},
		&calpb.Maxrequest{
			Number: 5,
		},
		&calpb.Maxrequest{
			Number: 3,
		},
		&calpb.Maxrequest{
			Number: 6,
		},
		&calpb.Maxrequest{
			Number: 2,
		},
		&calpb.Maxrequest{
			Number: 20,
		},
		&calpb.Maxrequest{
			Number: 0,
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
