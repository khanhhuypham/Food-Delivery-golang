package main

import (
	"Food-Delivery/proto-buffer/gen/calculatorpb"
	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	cc, err := grpc.Dial("localhost:5006", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("err while dial %v: ", err)
	}

	defer cc.Close()
	client := calculatorpb.NewCalculatorClient(cc)
	callSum(client)
	log.Println("Service client %f", client)
}

func callSum(c calculatorpb.CalculatorClient) {
	resp, err := c.Sum(context.Background(), &calculatorpb.SumRequest{
		Num1: 1, Num2: 2,
	})

	if err != nil {
		log.Fatalf("Sum called with err %v", err)
	}

	log.Println("Sum called with resp %v", resp.GetResult())
}
