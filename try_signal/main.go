package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	fmt.Println("started")

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		doSomethingLoop(ctx)
		wg.Done()
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigCh
		fmt.Println("signal received")
		cancel()
	}()

	wg.Wait()
	fmt.Println("finished")
}

func doSomethingLoop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			doSomething()
		}
	}
}

func doSomething() {
	fmt.Println("doSomething started.")
	time.Sleep(3 * time.Second)
	fmt.Println("doSomething finished.")
}