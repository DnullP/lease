package main

import (
	"context"
	"flag"
	"fmt"
	"lease/pb"
	"lease/utils"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type CenterNodeServer struct {
	pb.UnimplementedCenterNodeServer

	needWrite    bool
	writeChannel chan utils.Item
	data         map[string]string
	maxOutTime   int64
	// TODO maxOutTime need to align to every data
}

func NewCenterNodeServer() *CenterNodeServer {
	return &CenterNodeServer{
		data:         make(map[string]string),
		writeChannel: make(chan utils.Item, 10000),
		needWrite:    false,
		maxOutTime:   0,
	}
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
	log.Println("Lease: ", outTime)
	return response, nil
}

// TODO: need to use outdate to try to delete the data
func (s *CenterNodeServer) WriteData(ctx context.Context, req *pb.WriteDataRequest) (*pb.WriteDataResponse, error) {

	key, value := req.GetKey(), req.GetValue()

	s.writeChannel <- utils.Item{
		Key:   key,
		Value: value,
	}

	if !s.needWrite {
		s.needWrite = true
		go func() {
		LOOP:
			for {
				time.Sleep(100 * time.Millisecond)
				if s.maxOutTime < time.Now().Unix() {

					for {
						select {
						case data := <-s.writeChannel:
							s.data[data.Key] = data.Value
							fmt.Println("write data: ", data.Key, data.Value)
						default:
							s.needWrite = false
							break LOOP
						}
					}
				}
			}
		}()
	}
	return &pb.WriteDataResponse{State: 1}, nil
}

func main() {
	server := grpc.NewServer()

	centerServer := NewCenterNodeServer()

	pb.RegisterCenterNodeServer(server, centerServer)

	var port int
	flag.IntVar(&port, "port", 50051, "The server port")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("gRPC server is running on port: ", port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
