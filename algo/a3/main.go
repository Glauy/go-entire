package main

import "fmt"

type LinkedNode struct {
	Val  int
	Next *LinkedNode
}

func main() {
	head := &LinkedNode{}
	cur := head
	var next *LinkedNode
	var tmp *LinkedNode
	//cur获取下一个节点next
	//next获取下一个 tmp
	//next.Next = cur
	//cur = next
	//

	for cur != nil {
		next = head.Next
		tmp = next.Next
		next.Next = cur
		fmt.Println(cur.Val)
		cur = tmp
	}
}
