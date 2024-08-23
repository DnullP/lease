package main

import (
	"context"
	"fmt"
	"lease/pb"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type CenterNodeServer struct {
	pb.UnimplementedCenterNodeServer

	needWrite    bool
	writeChannel chan struct {
		key   string
		value string
	}
	data       map[string]string
	maxOutTime int64
}

func (s *CenterNodeServer) RequestData(ctx context.Context, req *pb.RequestDataRequest) (*pb.RequestDataResponse, error) {
	if s.needWrite {
		response := &pb.RequestDataResponse{
			Data:  s.data[req.GetDataName()],
			Lease: s.maxOutTime,
		}
		return response, nil
	}
	outTime := time.Now().Unix() + 10
	s.maxOutTime = outTime
	response := &pb.RequestDataResponse{
		Data:  s.data[req.GetDataName()],
		Lease: outTime,
	}
	return response, nil
}

func (s *CenterNodeServer) Write(key, value string) {

	s.writeChannel <- struct{ key, value string }{key, value}

	if !s.needWrite {
		{
			s.needWrite = true

			go func() {
				for {
					time.Sleep(100 * time.Millisecond)

					if s.maxOutTime < time.Now().Unix() {

						for data := range s.writeChannel {
							s.data[data.key] = data.data
						}
						s.needWrite = false
						break
					}
				}
			}()
		}
	}
}

func main() {
	server := grpc.NewServer()

	centerServer := &CenterNodeServer{}

	pb.RegisterCenterNodeServer(server, centerServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("gRPC server is running on port 50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	for {
		key := ""
		value := ""
		log.Println("Please input kv")
		fmt.Scan(&key)
		fmt.Scan(&value)
		centerServer.Write(key, value)
	}
}
