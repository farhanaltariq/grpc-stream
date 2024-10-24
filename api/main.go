package main

import (
	"context"
	"log"
	"net"

	pb "github.com/farhanaltariq/protomessager/go/protomessagergateway"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMapServiceServer
}

func (s *server) Map(ctx context.Context, req *pb.MapRequest) (*pb.MapResponse, error) {
	return &pb.MapResponse{Alt: req.Alt, Lat: req.Lat, Lon: req.Lon, Status: "OK"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMapServiceServer(s, &server{})

	log.Println("Server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
