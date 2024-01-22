package structure

type Node struct {
	Val  interface{}
	Next *Node
}

type LinkedList struct {
	Head *Node
}
