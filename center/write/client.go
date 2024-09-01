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
	var address string
	flag.StringVar(&address, "address", "localhost:50051", "address of the server")
	flag.Parse()

	var addr = flag.String("addr", address, "the address to connect to")

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
