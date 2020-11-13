package main

import (
	"context"
	"io"
	"log"
	"net"
	"time"

	"../libpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

var max = 0

func (*server) Adder(ctx context.Context, req *libpb.SumRequest) (*libpb.SumResponse, error) {
	n1 := req.GetNumber1()
	n2 := req.GetNumber2()
	sum := n1 + n2
	res := &libpb.SumResponse{
		Result: sum,
	}
	return res, nil
}

func (*server) Primes(req *libpb.PrimeRequest, stream libpb.Calculator_PrimesServer) error {
	number := req.GetNumber()
	k := 2
	for number > 1 {
		if int(number)%k == 0 {
			res := &libpb.PrimeResponse{
				Prime: int32(k),
			}
			stream.Send(res)
			number /= int32(k)
			time.Sleep(time.Second * 1)
		} else {
			k++
		}
	}
	return nil
}
func (*server) Average(stream libpb.Calculator_AverageServer) error {
	flag, sum := 0, 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			res := &libpb.AVGResponse{
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

func (*server) Max(stream libpb.Calculator_MaxServer) error {
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
			res := &libpb.MaxResponse{
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
	reflection.Register(s)
	libpb.RegisterCalculatorServer(s, &server{})
	if err != s.Serve(lis) {
		log.Fatalln(err)
	}

}
