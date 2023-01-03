package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/tamboto2000/learning-go-grpc/helloworld/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewHelloWorldClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "franklin"})
	if err != nil {
		log.Fatal("SayHello error:", err.Error())
	}

	log.Printf("Say hello to %s\n", r.Name)
}
