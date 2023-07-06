package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func consumer(queue chan int) {
	defer wg.Done()
	for {
		data := <-queue
		fmt.Println(data)
		time.Sleep(time.Second)
	}
}

func main() {
	var msg chan int
	msg = make(chan int)
	wg.Add(1)
	go consumer(msg)
	msg <- 1
	msg <- 2
	msg <- 3
	close(msg) // 关闭channel
	wg.Wait()
}
