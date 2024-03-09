package tree

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func New(val int) *Node {
	return &Node{Val: val}
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

func (n *Node) Search(val int) *Node {
	if n == nil {
		return nil
	}

	if val < n.Val {
		return n.Left.Search(val)
	} else if val > n.Val {
		return n.Right.Search(val)
	} else {
		return n
	}
}

func (n *Node) Insert(val int) {
	if val < n.Val {
		if n.Left == nil {
			n.Left = &Node{Val: val}
		} else {
			n.Left.Insert(val)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Val: val}
		} else {
			n.Right.Insert(val)
		}
	}
}

func (n *Node) IsEmpty() bool {
	return n == nil
}

func (n *Node) Height() int {
	if n == nil {
		return 0
	}
	return max(n.Left.Height(), n.Right.Height()) + 1
}

func (n *Node) Depth() int {
	if n == nil {
		return 0
	}
	return max(n.Left.Depth(), n.Right.Depth()) + 1
}

func (n *Node) CountLeaves() int {
	if n == nil {
		return 0
	}
	if n.Left == nil && n.Right == nil {
		return 1
	}
	return n.Left.CountLeaves() + n.Right.CountLeaves()
}

func (n *Node) Delete(val int) {
	if n == nil {
		return
	}

	if val < n.Val {
		n.Left.Delete(val)
	} else if val > n.Val {
		n.Right.Delete(val)
	} else {
		if n.Left == nil {
			n = n.Right
		} else if n.Right == nil {
			n = n.Left
		} else {
			minNode := n.Right
			for minNode.Left != nil {
				minNode = minNode.Left
			}
			n.Val = minNode.Val
			n.Right.Delete(minNode.Val)
		}
	}
}
