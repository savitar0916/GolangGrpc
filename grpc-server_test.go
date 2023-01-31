package main

import (
	pb "Guestbooks/proto"
	"context"
	"testing"
)

func Benchmark_GetGuestbook(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.GetGuestbook(context.Background(), &pb.GetGuestbookRequest{Query: "Exercise"})
	}
	b.StopTimer()
}
