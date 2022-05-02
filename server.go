package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/thesheps/glimr/proto"
	"google.golang.org/grpc"
)

var listenAddress = "localhost:1234"
var voteResponse pb.RequestVoteResponse

type glimrServer struct {
	pb.UnimplementedGlimrServer
}

func (gs *glimrServer) RequestVote(ctx context.Context, req *pb.RequestVoteRequest) (*pb.RequestVoteResponse, error) {
	voteResponse = pb.RequestVoteResponse{}

	return &voteResponse, nil
}

func main() {
	conn, err := net.Listen("tcp", listenAddress)

	if err != nil {
		log.Fatal("tcp connection error: ", err.Error())
	}

	grpcServer := grpc.NewServer()
	glimrServer := glimrServer{}
	pb.RegisterGlimrServer(grpcServer, &glimrServer)

	fmt.Println("Starting glimr server at: ", listenAddress)

	if err := grpcServer.Serve(conn); err != nil {
		log.Fatal(err.Error())
	}
}
