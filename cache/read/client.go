package main

import (
	"context"
	"flag"
	"fmt"
	"lease/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := flag.String("addr", "localhost:50052", "center node address")
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewCacheNodeClient(conn)

	for {
		var key string
		fmt.Scanln(&key)
		req := pb.ReadDataRequest{DataName: key}
		ctx := context.Background()
		response, err := client.ReadData(ctx, &req)
		if err != nil {
			panic(err)
		}
		if response.GetData() == "" {
			fmt.Println("No data")
			continue
		}
		fmt.Println(response.GetData())
	}
}
