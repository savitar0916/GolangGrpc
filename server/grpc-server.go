package main

import (
	pb "Guestbooks/proto"
	"context"
	"errors"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type GuestbooksServiceServer struct {
	pb.GuestbookServiceServer
}

func (g *GuestbooksServiceServer) GetGuestbook(ctx context.Context, request *pb.GetGuestbookRequest) (response *pb.GetGuestbookResponse, err error) {
	query := request.Query
	if query == "Exercise" {
		response = &pb.GetGuestbookResponse{
			Message: "Success",
			Status:  400,
			Guestbooks: []*pb.Guestbook{{
				Id:      "stwe123",
				Name:    "Adam",
				Title:   "Exercise",
				Content: "Take A Bicycle An Hour",
				Status:  true,
			}},
		}
		err = nil
		return
	}
	err = errors.New("query fail")
	return
}
func main() {
	port := ":8989"
	l, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("listen %s\n", port)
	s := grpc.NewServer()
	pb.RegisterGuestbookServiceServer(s, &GuestbooksServiceServer{})
	s.Serve(l)

}

// func compute(value int, id int, c chan int) {
// 	for i := 0; i < value; i++ {
// 		//fmt.Println("Received from channel", id)
// 		//fmt.Println(i)
// 		amt := time.Duration(rand.Intn(250))
// 		time.Sleep(time.Millisecond * amt)
// 		c <- i
// 	}

// 	fmt.Println("Channel", id, "is closing...")
// 	close(c)
// 	fmt.Println("Channel", id, "is closed.")

// }

// func main() {
// 	c1 := make(chan int, 100)
// 	c2 := make(chan int, 100)
// 	c3 := make(chan int, 100)
// 	c4 := make(chan int, 100)
// 	c5 := make(chan int, 100)

// 	go compute(100, 1, c1)
// 	go compute(100, 2, c2)
// 	go compute(100, 3, c3)
// 	go compute(100, 4, c4)
// 	go compute(100, 5, c5)

// 	var channels = []chan int{c1, c2, c3, c4, c5}

// 	for i := 0; i < 5; i++ {
// 		select {
// 		case v := <-channels[i]:
// 			fmt.Println("Received from channel", i+1, ":", v)
// 		default:
// 			fmt.Println("No value received from channel", i+1)
// 		}
// 	}
// 	var input string
// 	fmt.Scanln(&input)
// }
