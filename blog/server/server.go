package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"

	bookpb "../libpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var session *mgo.Session

type server struct{}
type bk struct {
	ID     bson.ObjectId `bson:"_id,omitempty"`
	Title  string
	Author string
	Price  float32
}

func (*server) CreateBook(ctx context.Context, req *bookpb.CreateBookRequest) (*bookpb.CreateBookResponse, error) {
	book := bk{
		Title:  req.GetBook().GetTitle(),
		Author: req.GetBook().GetAuthor(),
		Price:  req.GetBook().GetPrice(),
	}
	oid := bson.NewObjectId()
	book.ID = oid
	err := session.DB("bookstore").C("books").Insert(book)
	if err != nil {
		log.Fatalln(err)
	}
	res := &bookpb.CreateBookResponse{
		Id: oid.Hex(),
	}
	return res, nil
}

func (*server) ReadBook(ctx context.Context, req *bookpb.ReadBookRequest) (*bookpb.ReadBookResponse, error) {
	id := req.GetId()
	oid := bson.ObjectIdHex(id)
	book := bk{}
	err := session.DB("bookstore").C("books").Find(bson.M{"_id": oid}).One(&book)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	res := &bookpb.ReadBookResponse{
		Book: &bookpb.Book{
			Id:     book.ID.Hex(),
			Title:  book.Title,
			Author: book.Author,
			Price:  book.Price,
		},
	}
	return res, nil

}

func (*server) UpdateBook(ctx context.Context, req *bookpb.UpdateBookRequest) (*bookpb.UpdateBookResponse, error) {
	id := req.GetBook().GetId()

	oid := bson.ObjectIdHex(id)
	book := bk{
		ID:     oid,
		Title:  req.GetBook().GetTitle(),
		Author: req.GetBook().GetAuthor(),
		Price:  req.GetBook().GetPrice(),
	}
	err := session.DB("bookstore").C("books").UpdateId(oid, book)

	if err != nil {
		return nil, err
	}
	res := &bookpb.UpdateBookResponse{
		Id: oid.Hex(),
	}
	return res, nil
}

func (*server) DeleteBook(ctx context.Context, req *bookpb.DeleteBookRequest) (*bookpb.DeleteBookResponse, error) {
	id := req.GetId()
	if !bson.IsObjectIdHex(id) {
		err := errors.New("invalid id")
		return nil, err
	}
	oid := bson.ObjectIdHex(id)
	err := session.DB("bookstore").C("books").Remove(bson.M{"_id": oid})
	if err != nil {
		return nil, err
	}
	res := &bookpb.DeleteBookResponse{
		Status: true,
	}
	return res, nil

}

func (*server) ListBooks(req *bookpb.ListBooksRequest, stream bookpb.BookStore_ListBooksServer) error {
	bs := []bk{}

	err := session.DB("bookstore").C("books").Find(nil).All(&bs)
	if err != nil {
		return err
	}
	for _, v := range bs {
		res := &bookpb.ListBooksResponse{
			Book: &bookpb.Book{
				Id:     v.ID.Hex(),
				Title:  v.Title,
				Author: v.Author,
				Price:  v.Price,
			},
		}
		err := stream.Send(res)
		if err != nil {
			return err
		}
	}
	return nil

}

func init() {
	var err error
	session, err = mgo.Dial("mongodb://mahmoud:123456@localhost/bookstore")
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	reflection.Register(s)

	bookpb.RegisterBookStoreServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}

}
