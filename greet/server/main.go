package main

import (
	"log"
	"net"

	pb "github.com/leonardoas10/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "localhost:50051"

type Server struct {
	pb.GreetServiceServer
}

func main()  {
	list, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on:%v\n", err)
	}

	log.Printf("Listening on %v\n", addr)

	opts := []grpc.ServerOption{}
	tls := true // change that to false if needed

	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Failed loading certificates: %v\n", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	//opts = append(opts, grpc.ChainUnaryInterceptor(LogInterceptor(), CheckHeaderInterceptor()))

	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(list); err != nil {
		log.Fatalf("Failed to server: %v\n", err)
	}
}