package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"../libpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	c := libpb.NewGreetClient(conn)
	doUnaryWithDeadline(c, 3*time.Second)
	doUnaryWithDeadline(c, 5*time.Second)

}

func doUnaryWithDeadline(c libpb.GreetClient, t time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), t)
	req := &libpb.GreetingRequest{
		Greeting: &libpb.Greeting{
			Firstname: "Mahmoud",
			Lastname:  "Mekki",
		},
	}
	defer cancel()
	res, err := c.Greeting(ctx, req)
	if err != nil {
		statuserr, ok := status.FromError(err)
		if ok {
			if statuserr.Code() == codes.DeadlineExceeded {
				fmt.Println("Time out was hit,Deadline exceeded!!")
			} else {
				fmt.Println(statuserr)
			}

		} else {
			fmt.Println(err)
		}
		return
	}
	fmt.Println(res.Result)
}
