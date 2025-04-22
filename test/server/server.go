package main

import (
	"Food-Delivery/proto-buffer/gen/calculatorpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServer
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Println("sum called...")
	resp := &calculatorpb.SumResponse{
		Result: req.GetNum1() + req.GetNum2(),
	}

	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:5006")
	if err != nil {
		log.Fatal("error while listening on 0.0.0.0:5006")
	}
	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServer(s, &server{})

	fmt.Print("Calculator is running....")
	err = s.Serve(listen)

	if err != nil {
		log.Fatal("error while serving on localhost:5006")
	}

}
