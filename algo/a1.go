package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Pair struct {
	group int
	num   int
}

func Shuffle(origin []Pair) []Pair {
	println(origin)
	n := len(origin)
	source := rand.New(rand.NewSource(time.Now().Unix()))
	for i := range origin {
		j := i + source.Intn(n-i)
		origin[i], origin[j] = origin[j], origin[i]
	}
	return origin
}

func main() {

	groupList := [3][3]string{{"p11", "p12", "p13"}, {"p21", "p22", "p23"}, {"p31", "p32", "p33"}}
	//
	//bitMap := bitmap.New(2 * 3)
	//println(group_list)

	var origin []Pair
	for i, v := range groupList {
		for j, _ := range v {
			//fmt.Printf("arr[%v][%v]=%v \t \n", i, j, v2)
			origin = append(origin, Pair{i, j})
		}
	}

	shuffled := Shuffle(origin[:])
	//fmt.Println(shuffled)

	// (group - 1) * size = num
	cursor := 0
	nextCursor := 0
	for {
		if cursor >= len(shuffled)-1 {
			break
		}
		nextCursor = cursor + 1
		//用于探测是不是最后一个节点
		if nextCursor+1 == len(shuffled)-1 {
			fmt.Printf("(%s, %s, %s)\n", groupList[shuffled[cursor].group][shuffled[cursor].num],
				groupList[shuffled[nextCursor].group][shuffled[nextCursor].num],
				groupList[shuffled[nextCursor+1].group][shuffled[nextCursor+1].num])

		} else {
			fmt.Printf("(%s, %s)\n", groupList[shuffled[cursor].group][shuffled[cursor].num], groupList[shuffled[nextCursor].group][shuffled[nextCursor].num])
		}
		cursor = nextCursor + 1
	}

	fmt.Println("结束了")

}
