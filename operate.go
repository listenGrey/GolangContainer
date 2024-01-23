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

func Constructor(val interface{}) *LinkedList {
	var list LinkedList
	newNode := &Node{Val: val, Next: nil}
	if list.Head == nil {
		list.Head = newNode
	}
	return &list
}

func Print(list *LinkedList) {
	cur := list.Head
	for cur != nil && cur.Next != nil {
		fmt.Printf("%v -> ", cur.Val)
		cur = cur.Next
	}
	if cur.Next == nil {
		fmt.Printf("%v", cur.Val)
	}
}

func Split(list *LinkedList) (*LinkedList, *LinkedList) {
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

func SearchByValue(list *LinkedList, val interface{}, equals func(interface{}, interface{}) bool) (bool, error) {
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

func SearchByIndex(list *LinkedList, index int) (*Node, error) {
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

func Reverse(list *LinkedList) error {
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

func RemoveAtHead(list *LinkedList) error {
	if list.Head == nil {
		return errors.New("this is an empty list")
	}
	list.Head = list.Head.Next
	return nil
}

func RemoveAtTail(list *LinkedList) error {
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

func RemoveByIndex(list *LinkedList, index int) error {
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

func RemoveByValue(list *LinkedList, val interface{}, equals func(interface{}, interface{}) bool) error {
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

func Len(list *LinkedList) int {
	length := 0
	cur := list.Head

	for cur != nil {
		length++
		cur = cur.Next
	}

	return length
}

func InsertAtHead(list *LinkedList, val interface{}) {
	newNode := &Node{Val: val, Next: list.Head}
	list.Head = newNode
}

func InsertAtTail(list *LinkedList, val interface{}) {
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

func InsertAtIndex(list *LinkedList, val interface{}, index int) {
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

func DetectCycle(list *LinkedList) bool {
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

func DetectIntersection(list1, list2 *LinkedList) *Node {
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

func DetectMiddle(list *LinkedList) *Node {
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

func Clone(ori *LinkedList) *LinkedList {
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
