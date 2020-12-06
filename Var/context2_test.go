package Var

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestT(t *testing.T) {
	//cancel
	ctx, cancel := context.WithCancel(context.Background())
	go work(ctx, "work1")

	time.Sleep(time.Second * 3)
	cancel()
	time.Sleep(time.Second * 1)

	// with value
	ctx1, valueCancel := context.WithCancel(context.Background())
	valueCtx := context.WithValue(ctx1, "key", "test value context")
	go workWithValue(valueCtx, "value work", "key")
	time.Sleep(time.Second * 3)
	valueCancel()

	// timeout
	ctx2, timeCancel := context.WithTimeout(context.Background(), time.Second*3)
	go work(ctx2, "time cancel")
	time.Sleep(time.Second * 5)
	timeCancel()

	// deadline
	ctx3, deadlineCancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
	go work(ctx3, "deadline cancel")
	time.Sleep(time.Second * 5)
	deadlineCancel()

	time.Sleep(time.Second * 3)

}

func workWithValue(ctx context.Context, name string, key string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value(key))
			println(name, " get message to quit")
			return
		default:
			println(name, " is running", time.Now().String())
			time.Sleep(time.Second)
		}
	}
}

func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			println(name, " get message to quit")
			return
		default:
			println(name, " is running", time.Now().String())
			time.Sleep(time.Second)
		}
	}
}
