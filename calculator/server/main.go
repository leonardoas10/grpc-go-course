package main

import (
	"log"
	"net"

	pb "github.com/leonardoas10/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main()  {
	list, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on:%v\n", err)
	}

	log.Printf("Listening on %v\n", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err = s.Serve(list); err != nil {
		log.Fatalf("Failed to server: %v\n", err)
	}
}