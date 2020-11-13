package main

import (
	"context"
	"fmt"
	"log"

	bookpb "../libpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	c := bookpb.NewBookStoreClient(conn)

	id := doCreate(c)

	doRead(c, id)

}

func doCreate(c bookpb.BookStoreClient) string {
	req := &bookpb.CreateBookRequest{
		Book: &bookpb.Book{
			Title:  "Life with mekki",
			Author: "Mahmoud Mekki",
			Price:  52.56,
		},
	}
	res, err := c.CreateBook(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res.Id)
	return res.Id
}

func doRead(c bookpb.BookStoreClient, id string) {
	req := &bookpb.ReadBookRequest{
		Id: id,
	}
	res, err := c.ReadBook(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}
