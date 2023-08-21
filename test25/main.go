package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	Sleep(2)
	fmt.Println("done")
}

func Sleep(seconds int) {
	<-time.After(time.Second * time.Duration(seconds))
}
