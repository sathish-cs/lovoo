package main

import (
	"context"
	"log"
	"net"

	pb "calculator-gRPC-golang/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	protocol = "tcp"
	port     = ":50051"
)

// Create  struct to implement interface 
type server struct{}

// Add function takes context from request and outputs
func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddReply, error) {
	return &pb.AddReply{N1: in.N1 + in.N2}, nil
}
// Subtract function takes context from request and outputs
func (s *server) Subtract(ctx context.Context, in *pb.SubtractRequest) (*pb.SubtractReply, error) {
	return &pb.SubtractReply{N1: in.N1 - in.N2}, nil
}
// Multiply function takes context from request and outputs
func (s *server) Multiply(ctx context.Context, in *pb.MultiplyRequest) (*pb.MultiplyReply, error) {
	return &pb.MultiplyReply{N1: in.N1 * in.N2}, nil
}
// Divide function takes context from request and outputs
func (s *server) Divide(ctx context.Context, in *pb.DivideRequest) (*pb.DivideReply, error) {
	return &pb.DivideReply{N1: in.N1 / in.N2}, nil
}

func main() {
	// TCP listener on port 50051
	lis, err := net.Listen(protocol, port)
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}
	// Create function from grpc to create server object
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
