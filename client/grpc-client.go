package main

import (
	pb "Guestbooks/proto"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8989", grpc.WithInsecure())
	if err != nil {
		fmt.Println("dial error: " + err.Error())
	}

	defer conn.Close()
	client := pb.NewGuestbookServiceClient(conn)

	request := new(pb.GetGuestbookRequest)
	request.Query = "Exercise"
	response, err := client.GetGuestbook(context.Background(), request)

	if err != nil {
		fmt.Println("response error" + err.Error())
	}

	fmt.Println(response)
}
