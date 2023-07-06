package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("hello from goroutine %d\n", i)
		}()
		//go func(i int) {
		//	fmt.Printf("hello from goroutine %d\n", i)
		//}(i)
	}
	time.Sleep(time.Minute)
}
