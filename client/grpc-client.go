package main

import (
	pb "Guestbooks/proto"
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:8989"
	users   = 2000
	request = 100
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("did not connect:", err)
	}
	defer conn.Close()
	client := pb.NewGuestbookServiceClient(conn)
	fmt.Println("------------------------")
	fmt.Printf("%v個用戶\n", users)
	fmt.Println("------------------------")
	fmt.Printf("%v個請求\n", request)
	fmt.Println("-----開始發送Request-----")
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(users)
	for i := 0; i < users; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < request; j++ {
				_, err := client.GetGuestbook(context.Background(), &pb.GetGuestbookRequest{Query: "Exercise"})
				if err != nil {
					fmt.Printf("could not greet: %v\n", err)
				}
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			}
		}()
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("------------------------")
	fmt.Println("花費時間:", elapsed)
	fmt.Println("------------------------")
	fmt.Println("吞吐量為:", float64(users*request)/elapsed.Seconds())
	// fmt.Printf("Time elapsed: %s\n", elapsed)
	// fmt.Printf("Throughput: %f requests/second\n", float64(N*M)/elapsed.Seconds())
}

// var wg sync.WaitGroup
// var totalTime time.Duration

// func sendRequest(client pb.GuestbookServiceClient) {

// 	defer wg.Done()

// 	response, err := client.GetGuestbook(context.Background(), &pb.GetGuestbookRequest{Query: "Exercise"})
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	} else {
// 		fmt.Println(response)
// 	}

// }

// func main() {
// 	start := time.Now()
// 	conn, err := grpc.Dial("localhost:8989", grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	defer conn.Close()
// 	client := pb.NewGuestbookServiceClient(conn)
// 	numOfRoutines := 2000
// 	requestCount := 20

// 	wg.Add(numOfRoutines * requestCount)
// 	for i := 0; i < numOfRoutines; i++ {
// 		for j := 0; j < requestCount; j++ {
// 			go sendRequest(client)
// 		}
// 	}
// 	wg.Wait()
// 	elapsed := time.Since(start)
// 	totalTime += elapsed
// 	fmt.Printf("%v個用戶\n", numOfRoutines)
// 	fmt.Println("-----------------------")
// 	fmt.Printf("%v個請求\n", requestCount)
// 	fmt.Println("-----------------------")
// 	fmt.Println("花費時間:", totalTime)
// 	fmt.Println("-----------------------")
// 	fmt.Println("吞吐量為:", (totalTime / time.Duration(numOfRoutines*requestCount)).String())

// }
