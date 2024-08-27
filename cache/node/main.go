package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "lease/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DatawithDDL struct {
	data string
	ddl  int64
}

// CacheNodeServer is the server for CacheNode service
type CacheNodeServer struct {
	pb.UnimplementedCacheNodeServer

	data map[string]DatawithDDL
	conn *grpc.ClientConn
}

func NewCacheNodeServer() *CacheNodeServer {
	addr := flag.String("addr", "localhost:50051", "the address to connect to")
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	return &CacheNodeServer{
		data: make(map[string]DatawithDDL),
		conn: conn,
	}
}

// OutdateData handles the OutdateData RPC
func (s *CacheNodeServer) OutdateData(ctx context.Context, req *pb.OutdateDataRequest) (*pb.OutdateDataResponse, error) {
	dataName := req.GetDataName()
	delete(s.data, dataName)

	// log the nowdata
	fmt.Printf("outdate the data: %s\n", dataName)

	return &pb.OutdateDataResponse{
		Success: true,
	}, nil
}

func (s *CacheNodeServer) ReadData(ctx context.Context, req *pb.ReadDataRequest) (*pb.ReadDataResponse, error) {
	dataName := req.GetDataName()
	data, ok := s.data[dataName]

	RequestCenterData := func() string {
		client := pb.NewCenterNodeClient(s.conn)
		requestToCenter := &pb.RequestDataRequest{
			DataName: dataName,
		}
		response, err := client.RequestData(context.Background(), requestToCenter)
		if err != nil {
			log.Println("Failed to request data: %v", err)
			return "No data"
		}
		newData := response.GetData()
		newDDL := response.GetLease()
		s.data[dataName] = DatawithDDL{
			data: newData,
			ddl:  newDDL,
		}
		return newData
	}

	if !ok {
		newData := RequestCenterData()

		return &pb.ReadDataResponse{
			Data: newData,
		}, nil
	}
	if data.ddl < time.Now().Unix() {
		newData := RequestCenterData()
		return &pb.ReadDataResponse{
			Data: newData,
		}, nil
	}

	return &pb.ReadDataResponse{
		Data: data.data,
	}, nil
}

func main() {
	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register CacheNode and CenterNode services
	pb.RegisterCacheNodeServer(grpcServer, NewCacheNodeServer())

	// Listen on a port
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Start the server
	fmt.Println("gRPC server is running on port 50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	for {
		query := ""
		fmt.Scan(query)

	}
}
