package main

import "fmt"

func main() {
	ForR()
}

func ForI() {

}

func ForR() {
	arr := []int{9, 8, 7, 6}
	for index, value := range arr {
		fmt.Printf("%d => %d\n", index, value)
	}
}
