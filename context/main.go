package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	// ...
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled")
		return
	case <-time.After(5 * time.Second):
		fmt.Println("Hotel booking done")
	}
}
