package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"os"
	pb "proxy/client/proxy"
)

const (
	serviceAddress = "localhost:9999"
)

func main()  {
	if len(os.Args) < 2 {
		fmt.Println("no url provided")
		return
	}
	url := pb.Url{Value: os.Args[1]}

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

	r, err := c.Get(context.Background(), &url)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Result: \n%s", r.Value)
}


