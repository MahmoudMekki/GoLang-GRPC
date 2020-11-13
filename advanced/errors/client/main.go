package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/status"

	lipb "../libpb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	c := lipb.NewSquareRootClient(conn)
	req := &lipb.SquareRootRequest{
		Number: 0,
	}
	res, err := c.SquareRoot(context.Background(), req)
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())
			return
		}
		fmt.Println(err)
	}

	fmt.Println(res)

}
