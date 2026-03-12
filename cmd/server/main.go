package main

import (
	"log"
	"net"

	newsv1 "github.com/rbalusup/news-grpc/api/news/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthv1 "google.golang.org/grpc/health/grpc_health_v1"

	ingrpc "github.com/rbalusup/news-grpc/internal/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	newsv1.RegisterNewsServiceServer(srv, ingrpc.NewServer())
	healthSrv := health.NewServer()
	healthv1.RegisterHealthServer(srv, healthSrv)

	log.Printf("gRPC server is running on %s", lis.Addr().String())
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
