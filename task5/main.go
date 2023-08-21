package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/

func main() {
	DoJob(1)
}

func DoJob(seconds uint) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Duration(seconds)*time.Second)
	defer cancel()

	dataChan := make(chan string)
	runWorker(ctxWithTimeout, dataChan)

	go func() {
		for i := 0; i < 100000; i++ {
			dataChan <- "wrote in channel " + strconv.Itoa(rand.Int())
			time.Sleep(1 * time.Microsecond)
		}
	}()

	time.Sleep(time.Duration(seconds) * (time.Second + 1))
}

func runWorker(ctx context.Context, dataChan <-chan string) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("worker : DONE\n")
				return
			case data := <-dataChan:
				fmt.Printf("worker data: %s\n", data)
			}
		}
	}()
}
