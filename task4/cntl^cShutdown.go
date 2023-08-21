package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают
произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

func main() {
	DoJob()
}

func DoJob() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	sigTerm := make(chan os.Signal)
	signal.Notify(sigTerm, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigTerm
		cleanup()
		os.Exit(1)
	}()

	dataChan := make(chan string)
	runWorkers(sigTerm, dataChan, 4)

	go func() {
		for i := 0; i < 100000; i++ {
			dataChan <- "write in channel " + strconv.Itoa(rand.Int())
			time.Sleep(1 * time.Microsecond)
		}
	}()

	wg.Wait()
}

func cleanup() {
	fmt.Println("stop")
}

func runWorkers(done <-chan os.Signal, dataChan <-chan string, workersNum int) {
	for i := 0; i < workersNum; i++ {
		go func(workerNum int) {
			for {
				select {
				case <-done:
					fmt.Printf("worker number: %d DONE\n", workerNum)
					return
				case data := <-dataChan:
					fmt.Printf("worker number: %d, data: %s\n", workerNum, data)
				}
			}
		}(i)
	}
}
