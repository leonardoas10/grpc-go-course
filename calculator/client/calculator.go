package main

import (
	"context"
	"log"

	pb "github.com/leonardoas10/grpc-go-course/calculator/proto"
)

func doCalculator(c pb.CalculatorServiceClient)  {
	log.Println("doGreetr was invoked")
	res, err := c.Calculator(context.Background(), &pb.CalculatorRequest{
		NumberOne: 10,
		NumberTwo: 3,
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %v\n", res.Result)
}