package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "github.com/Henate/Streaming-gRPC/proto"
)

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

const PORT = "9001"

func main() {
	server := grpc.NewServer()	//创建server 端的抽象对象
	pb.RegisterSearchServiceServer(server, &SearchService{})	//注册Server

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server.Serve(lis)
}