package linkedList

import (
	"errors"
	"fmt"
)

type Node struct {
	Val  interface{}
	Next *Node
}

type LinkedList struct {
	Head *Node
}

type cloneList struct {
	oriList *LinkedList
	resList *LinkedList
}

func New() *LinkedList {
	var list LinkedList
	return &list
}

func (list *LinkedList) InsertAtHead(val interface{}) {
	newNode := &Node{Val: val, Next: list.Head}
	list.Head = newNode
}

func (list *LinkedList) InsertAtTail(val interface{}) {
	newNode := &Node{Val: val, Next: nil}
	if list.Head == nil { // if input linked list is null, new node is head node
		list.Head = newNode
		return
	}
	cur := list.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = newNode
}

func (list *LinkedList) InsertAtIndex(val interface{}, index int) {
	newNode := &Node{Val: val, Next: nil}

	if index <= 0 {
		newNode.Next = list.Head
		list.Head = newNode
		return
	}

	cur := list.Head
	for i := 0; i < index-1 && cur != nil; i++ {
		cur = cur.Next
	}

	if cur == nil {
		for cur = list.Head; cur.Next != nil; cur = cur.Next {
		}
		cur.Next = newNode
		return
	}

	newNode.Next = cur.Next
	cur.Next = newNode
}

func (list *LinkedList) RemoveAtHead() error {
	if list.Head == nil {
		return errors.New("this is an empty list")
	}
	list.Head = list.Head.Next
	return nil
}

func (list *LinkedList) RemoveAtTail() error {
	if list.Head == nil {
		return errors.New("this is an empty list")
	}
	if list.Head.Next == nil {
		list.Head = nil
		return nil
	}
	cur := list.Head
	for cur.Next.Next != nil {
		cur = cur.Next
	}
	cur.Next = nil
	return nil
}

func (list *LinkedList) RemoveByIndex(index int) error {
	if index < 0 {
		return errors.New("index must >= 0")
	}
	if list.Head == nil {
		return errors.New("this is an empty list")
	}

	if index == 0 {
		list.Head = list.Head.Next
		return nil
	}

	cur := list.Head
	for i := 0; i < index-1 && cur.Next != nil; i++ {
		cur = cur.Next
	}

	if cur.Next == nil {
		return errors.New("index out of range")
	}
	cur.Next = cur.Next.Next
	return nil
}

func (list *LinkedList) RemoveByValue(val interface{}, equals func(interface{}, interface{}) bool) error {
	if equals == nil {
		return errors.New("need a comparative function")
	}

	if list.Head == nil {
		return errors.New("this is an empty list")
	}

	for list.Head != nil && equals(list.Head.Val, val) {
		list.Head = list.Head.Next
	}

	cur := list.Head
	for cur != nil && cur.Next != nil {
		if equals(cur.Next.Val, val) {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return nil
}

func (list *LinkedList) MoveToHead(index int) error {
	if list == nil || list.Head == nil {
		return errors.New("this is an empty list")
	}
	if index < 0 {
		return errors.New("need correct index")
	}
	cur := list.Head
	var pre *Node
	for i := 0; i < index && cur != nil; i++ {
		pre = cur
		cur = cur.Next
	}

	if cur != nil {
		if pre != nil {
			pre.Next = cur.Next
			cur.Next = list.Head
			list.Head = cur
		}
		return nil
	}
	return errors.New("index out of range")
}

func (list *LinkedList) MoveToTail(index int) error {
	if list == nil || list.Head == nil {
		return errors.New("this is an empty list")
	}
	if index < 0 {
		return errors.New("need correct index")
	}
	cur := list.Head
	var pre *Node
	for i := 0; i < index && cur != nil; i++ {
		pre = cur
		cur = cur.Next
	}

	if cur != nil {
		if cur.Next != nil {
			pre.Next = cur.Next
			cur.Next = nil

			tail := list.Head
			for tail.Next != nil {
				tail = tail.Next
			}
			tail.Next = cur
		}
		return nil
	}
	return errors.New("index out of range")
}

func (list *LinkedList) SearchByValue(val interface{}, equals func(interface{}, interface{}) bool) (bool, error) {
	if equals == nil {
		return false, errors.New("need a comparative function")
	}

	cur := list.Head
	for cur != nil {
		if equals(cur.Val, val) {
			return true, nil
		}
		cur = cur.Next
	}
	return false, nil
}

func (list *LinkedList) SearchByIndex(index int) (*Node, error) {
	if index < 0 {
		return nil, errors.New("index must >= 0")
	}
	cur := list.Head
	for i := 0; cur != nil && i < index; i++ {
		cur = cur.Next
	}

	if cur == nil {
		return nil, errors.New("index out of range")
	}

	return cur, nil
}

func (list *LinkedList) Print() {
	cur := list.Head
	for cur != nil && cur.Next != nil {
		fmt.Printf("%v -> ", cur.Val)
		cur = cur.Next
	}
	if cur.Next == nil {
		fmt.Printf("%v", cur.Val)
	}
}

func (list *LinkedList) Split() (*LinkedList, *LinkedList) {
	if list.Head == nil || list.Head.Next == nil {
		return list, nil
	}

	slow := list.Head
	fast := list.Head

	var preList *Node

	for fast != nil && fast.Next != nil {
		preList = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	if preList != nil {
		preList.Next = nil
	}

	list1 := &LinkedList{Head: list.Head}
	list2 := &LinkedList{Head: slow}

	return list1, list2
}

func (list *LinkedList) Reverse() error {
	if list == nil {
		return errors.New("this is an empty list")
	}
	var pre *Node = nil
	cur := list.Head
	var next *Node = nil

	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	list.Head = pre
	return nil
}

func (list *LinkedList) Len() int {
	length := 0
	cur := list.Head

	for cur != nil {
		length++
		cur = cur.Next
	}

	return length
}

func (ori *LinkedList) Clone() *LinkedList {
	if ori == nil || ori.Head == nil {
		return nil
	}

	clone := &cloneList{
		oriList: ori,
		resList: &LinkedList{Head: nil},
	}

	oriCur := ori.Head
	var resCur *Node

	for oriCur != nil {
		resNode := &Node{Val: oriCur.Val, Next: nil}

		if clone.resList.Head == nil {
			clone.resList.Head = resNode
			resCur = clone.resList.Head
		} else {
			resCur.Next = resNode
			resCur = resNode
		}

		oriCur = oriCur.Next
	}

	return clone.resList
}

func (list *LinkedList) DetectCycle() bool {
	if list.Head == nil {
		return false
	}

	slow := list.Head
	fast := list.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}

func (list1 *LinkedList) DetectIntersection(list2 *LinkedList) *Node {
	if list1 == nil || list1.Head == nil || list2 == nil || list2.Head == nil {
		return nil
	}
	length1 := 0
	cur1 := list1.Head
	for cur1 != nil {
		length1++
		cur1 = cur1.Next
	}

	length2 := 0
	cur2 := list2.Head
	for cur2 != nil {
		length2++
		cur2 = cur2.Next
	}

	cur1 = list1.Head
	cur2 = list2.Head

	dec := 0

	if length1 > length2 {
		dec = length1 - length2
		for i := 0; i < dec; i++ {
			cur1 = cur1.Next
		}
	} else if length1 < length2 {
		dec = length2 - length1
		for i := 0; i < dec; i++ {
			cur2 = cur2.Next
		}
	}

	for cur1 != nil && cur2 != nil {
		if cur1 == cur2 {
			return cur1
		}
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	return nil
}

func (list *LinkedList) DetectMiddle() *Node {
	if list.Head == nil {
		return nil
	}

	slow := list.Head
	fast := list.Head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
