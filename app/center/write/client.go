package main

import (
	"context"
	"flag"
	"fmt"
	"lease/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "center node address")
)

func main() {
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewCenterNodeClient(conn)

	var key string
	var value string
	for {
		fmt.Scanln(&key)
		fmt.Scanln(&value)

		req := pb.WriteDataRequest{Key: key, Value: value}
		ctx := context.Background()
		respose, err := client.WriteData(ctx, &req)
		if err != nil {
			panic(err)
		}
		fmt.Println(respose.GetState())
	}

}
