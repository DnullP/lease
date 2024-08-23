package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "lease/pb"
	"lease/storage"

	"google.golang.org/grpc"
)

// CacheNodeServer is the server for CacheNode service
type CacheNodeServer struct {
	pb.UnimplementedCacheNodeServer
}

// OutdateData handles the OutdateData RPC
func (s *CacheNodeServer) OutdateData(ctx context.Context, req *pb.OutdateDataRequest) (*pb.OutdateDataResponse, error) {
	data_name := req.GetDataName()
	delete(storage.NodeCache, data_name)

	// log the nowdata
	fmt.Printf("outdate the data: %s\n", data_name)

	return &pb.OutdateDataResponse{
		Success: true,
	}, nil
}

// connect to the center node
// func connectToCenterNode() {

// }

func main() {
	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register CacheNode and CenterNode services
	pb.RegisterCacheNodeServer(grpcServer, &CacheNodeServer{})

	// Listen on a port
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Start the server
	fmt.Println("gRPC server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	for {
		query := ""
		fmt.Scan(query)

	}
}
