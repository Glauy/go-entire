package main

import (
	"fmt"
	"math/rand"
)

//rand.New 新建rand
func random() {
	//创建一个Rand
	const seed int64 = 46
	source := rand.NewSource(seed) //seed相同的情况下，随机数生成器产生的数列是相同的
	rander := rand.New(source)
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", rander.Intn(100))
	}
	fmt.Println()
	source.Seed(seed + 10) //必须重置一下Seed
	rander2 := rand.New(source)
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", rander2.Intn(100))
	}
	fmt.Println()
}

//使用全局rand
func globalRand() {
	//使用全局Rand
	rand.Seed(1)                //如果对两次运行没有一致性要求，可以不设seed
	fmt.Println(rand.Int())     //随机生成一个整数
	fmt.Println(rand.Float32()) //随机生成一个浮点数
	fmt.Println(rand.Intn(10))  //10以内的随机整数，[0,10)
	fmt.Println(rand.Perm(10))  //把[0,10)上的整数随机打乱
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Shuffle(len(arr), func(i, j int) { //随机打乱一个给定的slice
		arr[i], arr[j] = arr[j], arr[i]
	})
	fmt.Println(arr)
}

func main() {
	//random()
	globalRand()
}
