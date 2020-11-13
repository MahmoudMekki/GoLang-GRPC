package main

import (
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	"../calpb"
)

type server struct{}

func (*server) Average(stream calpb.Average_AverageServer) error {
	flag, sum := 0, 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			res := &calpb.AVGResponse{
				Average: (float32(sum) / float32(flag)),
			}
			return stream.SendAndClose(res)

		}
		if err == nil {
			flag++
			sum += int(req.GetNumber())
		} else {
			break
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	calpb.RegisterAverageServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
