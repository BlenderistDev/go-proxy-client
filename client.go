package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "proxy/client/proxy"
)

const (
	serviceAddress = "localhost:9999"
)

func main()  {
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(conn)

	c := pb.NewProxyClient(conn)
	url := pb.Url{Value: "https://google.com"}

	r, err := c.Get(context.Background(), &url)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Result: \n%s", r.Value)
}


