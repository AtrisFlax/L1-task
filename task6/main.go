package main

import (
	"context"
	"fmt"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

// close channel
func main() {
	closeOutside()
	closeUsingChanStruct()
	closeUsingContext()
}

func closeUsingContext() {
	fmt.Println("closeUsingContext")
	channel := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				channel <- struct{}{}
				fmt.Println("closeUsingContext: channel closed")
				return
			default:
				fmt.Println("working...")
				time.Sleep(1 * time.Second)
			}

		}
	}(ctx)

	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	<-channel
	fmt.Println("Finish")
}

func closeUsingChanStruct() {
	fmt.Println("closeUsingChanStruct")
	channel := make(chan string)
	done := make(chan struct{})
	go func() {
		for {
			select {
			case channel <- "foo":
			case <-done:
				close(channel)
				fmt.Println("closeUsingChanStruct: channel closed")
				return
			}
		}
	}()

	go func() {
		time.Sleep(2 * time.Second)
		done <- struct{}{}
	}()

	/*for i := range channel {
		fmt.Println("reading from channel: ", i)
	}*/

	fmt.Println("Finish")
}

func closeOutside() {
	fmt.Println("closeOutside")
	channel := make(chan string)
	go func() {
		for {
			v, ok := <-channel
			if !ok {
				fmt.Println("closeOutside: channel closed")
				return
			}
			fmt.Println(v)
		}
	}()

	channel <- "data1"
	channel <- "data2"
	close(channel)
	time.Sleep(time.Second)
}
