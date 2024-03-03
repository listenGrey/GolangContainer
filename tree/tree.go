package tree

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (n *Node) PrintFront() {
	if n != nil {
		fmt.Printf("%d", n.Val)
		n.Left.PrintFront()
		n.Right.PrintFront()
	}
}

func (n *Node) PrintMid() {
	if n != nil {
		n.Left.PrintMid()
		fmt.Printf("%d", n.Val)
		n.Right.PrintMid()
	}
}

func (n *Node) PrintBack() {
	if n != nil {
		n.Left.PrintBack()
		n.Right.PrintBack()
		fmt.Printf("%d", n.Val)
	}
}
