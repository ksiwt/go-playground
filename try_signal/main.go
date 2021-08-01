package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	inputs := []string{}
	for i := 1; i <= 20; i++ {
		inputs = append(inputs, strconv.Itoa(i))
	}

	eg := errgroup.Group{}
	eg.Go(func() error {
		return doSomethingConcurrent(ctx, inputs)
	})

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigCh
		fmt.Println("signal received")
		cancel()
	}()

	if err := eg.Wait(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("completed")
}

func doSomethingConcurrent(ctx context.Context, inputs []string) error {
	allResource := int64(3)
	doSomethingResource := int64(1)

	sem := semaphore.NewWeighted(allResource)
	var wg sync.WaitGroup
	for _, v := range inputs {
		v := v
		if err := sem.Acquire(ctx, doSomethingResource); err != nil {
			wg.Wait()
			return err
		}
		go func() {
			wg.Add(1)
			doSomething(v)
			sem.Release(doSomethingResource)
			wg.Done()
		}()
	}
	wg.Wait()
	return nil
}

func doSomething(input string) {
	fmt.Println("doSomething started: " + input)
	time.Sleep(3 * time.Second)
	fmt.Println("doSomething finished: " + input)
}