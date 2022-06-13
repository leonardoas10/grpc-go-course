package main

import (
	"context"
	"log"

	pb "github.com/leonardoas10/grpc-go-course/greet/proto"
)

func doGreet(c pb.GreetServiceClient)  {
	log.Println("doGreetr was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Clemente",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %v\n", res.Result)
}