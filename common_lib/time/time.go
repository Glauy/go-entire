package main

import (
	"math/rand"
	"time"
)

func main() {

	// 睡眠100毫秒
	select {
	case <-time.NewTimer(100 * time.Millisecond).C:
	}

	// tips: 随机睡眠(0~1s)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}
