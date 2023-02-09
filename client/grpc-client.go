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
	fmt.Println("------------------------")
	fmt.Println("開始發送請求")
	fmt.Println("------------------------")
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(users)
	for i := 1; i <= users; i++ {
		go func() {
			defer wg.Done()
			for j := 1; j <= request; j++ {
				_, err := client.GetGuestbook(context.Background(), &pb.GetGuestbookRequest{Query: "Exercise"})
				if err != nil {
					fmt.Printf("Error: %v\n", err)
				}
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
				//fmt.Println(response)
			}
		}()
		if (i*request)%10000 == 0 {
			fmt.Println("目前請求數:", i*request)
		}
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("------------------------")
	fmt.Println("總共請求:", users*request)
	fmt.Println("------------------------")
	fmt.Println("總共時間:", elapsed)
	fmt.Println("------------------------")
	var RPS = fmt.Sprintf("%.2f", float64(users*request)/elapsed.Seconds())
	fmt.Println("吞吐量為:", RPS)
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
