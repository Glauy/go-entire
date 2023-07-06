package main

func main() {

}

// Node Invalid recursive type 'Node'
type Node struct {
	//自引用只能使用指针
	//left Node
	//right Node

	left  *Node
	right *Node

	// 这个也会报错
	// nn NodeNode
}

type NodeNode struct {
	node Node
}
