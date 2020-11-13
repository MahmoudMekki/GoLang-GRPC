package main

import (
	"io"
	"log"
	"net"

	"../calpb"
	"google.golang.org/grpc"
)

type server struct{}

var max = 0

func (*server) Max(stream calpb.Max_MaxServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalln(err)
			return err
		}
		number := req.GetNumber()
		if int(number) > max {
			max = int(number)
			res := &calpb.Maxresponse{
				Result: int32(max),
			}
			err := stream.Send(res)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	calpb.RegisterMaxServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}

}
