package main

import (
	"fmt"
	"sync"
	"time"
)

func isTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	for {
		select {
		case wg.Wait():
		case time.Sleep(timeout):
		}
	}
}

func main() {
	// 批量下载图片
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			// 假设下载图片的时间
			fmt.Println("Sleep: ", num, " second")
			time.Sleep(time.Duration(num) * time.Second)
		}(i)
	}
	// 设置超时时间
	if isTimeout(&wg, time.Second*5) {
		fmt.Println("Timeout")
		return
	}
	fmt.Println("Done")
	time.Sleep(time.Second * 10)
}
