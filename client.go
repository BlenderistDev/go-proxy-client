package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "proxy/client/proxy"
	"time"
)
func main()  {
	conn, err := grpc.Dial("localhost:9999", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	c := pb.NewProxyClient(conn)


	url := pb.Url{Value: "https://google.com"}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, err := c.Get(ctx, &url)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(r)
}


