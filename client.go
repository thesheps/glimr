package main

import (
	"context"
	"log"
	"time"

	pb "github.com/thesheps/glimr/proto"
	"google.golang.org/grpc"
)

func startClient() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())

	if err != nil {
		log.Fatal("tcp connection error: ", err.Error())
	}

	defer conn.Close()

	client := pb.NewGlimrClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	response, err := client.RequestVote(ctx, &pb.RequestVoteRequest{})

	if err != nil {
		log.Fatal("Error Requesting vote: ", err.Error())
	}

	log.Println(response)
}
