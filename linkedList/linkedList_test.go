package linkedList

import (
	"io/ioutil"
	"math/rand"
	"os"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	mock := &LinkedList{Head: nil}
	res := New()

	if reflect.TypeOf(mock) != reflect.TypeOf(res) {
		t.Errorf("New function is invalid")
	}
}

func TestLinkedList_Print(t *testing.T) {
	res := &LinkedList{}
	res.InsertAtTail(1)
	res.InsertAtTail(2)
	res.InsertAtTail(3)
	res.InsertAtTail(4)
	res.InsertAtTail(5)
	out := captureOutput(func() {
		res.Print()
	})
	if out != "1 -> 2 -> 3 -> 4 -> 5" {
		t.Errorf("Expected output: 1 -> 2 -> 3 -> 4 -> 5, got: %s", out)
	}
}

func captureOutput(f func()) string {
	// 将标准输出重定向到缓冲区
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// 调用函数
	f()

	// 恢复标准输出
	w.Close()
	os.Stdout = old

	// 读取缓冲区中的内容
	out, _ := ioutil.ReadAll(r)
	return string(out)
}

func TestLinkedList_InsertAtHead(t *testing.T) {
	value := rand.Int()
	res := New()
	res.InsertAtHead(value)

	if value != res.Head.Val {
		t.Errorf("Insert at head function is invalid")
	}
}

func TestLinkedList_InsertAtTail(t *testing.T) {
	value := rand.Int()
	res := New()
	res.InsertAtTail(value)
	res.InsertAtTail(value + 1)
	res.InsertAtTail(value + 2)
	res.InsertAtTail(value + 3)
	res.InsertAtTail(value + 4)
	res.InsertAtTail(value + 5)

	cur := res.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	if value+5 != cur.Val {
		t.Errorf("Insert at tail function is invalid")
	}
}

func TestLinkedList_InsertAtIndex(t *testing.T) {
	value := rand.Int()

	//case 1: index < 0, insert at head
	res1 := New()
	res1.InsertAtIndex(value, -1)
	if value != res1.Head.Val {
		t.Errorf("When the index < 0, insert at index function is invalid")
	}

	//case 2: 0 < index < end, insert at index
	res2 := New()
	res2.InsertAtTail(value)
	res2.InsertAtTail(value + 1)
	res2.InsertAtTail(value + 2)
	res2.InsertAtTail(value + 3)
	res2.InsertAtTail(value + 4)
	res2.InsertAtIndex(value, 2)
	if value != res2.Head.Next.Next.Val {
		t.Errorf("When the 0 < index < end, insert at index function is invalid")
	}

	//case 3: end < index, insert at end
	res3 := New()
	res3.InsertAtTail(value)
	res3.InsertAtTail(value + 1)
	res3.InsertAtTail(value + 2)
	res3.InsertAtIndex(value, 5)
	if value != res3.Head.Next.Next.Next.Val {
		t.Errorf("When the end < index, insert at index function is invalid")
	}
}

func TestLinkedList_RemoveAtHead(t *testing.T) {
	value := rand.Int()
	//case 1: head is nil
	res1 := New()
	err := res1.RemoveAtHead()
	if err == nil {
		t.Errorf("Remove at head function is invalid")
	}
	if res1.Head != nil {
		t.Errorf("Remove at head function is invalid")
	}

	//case 2: head is not nil
	res2 := New()
	res2.InsertAtTail(value)
	res2.InsertAtTail(value + 1)
	res2.InsertAtTail(value + 2)
	err = res2.RemoveAtHead()
	if err != nil {
		t.Errorf("%v", err)
	}
	if value == res2.Head.Val {
		t.Errorf("Remove at head function is invalid")
	}
}

func TestLinkedList_RemoveAtTail(t *testing.T) {
	value := rand.Int()
	//case 1: list is nil
	res1 := New()
	err := res1.RemoveAtTail()
	if err == nil {
		t.Errorf("Remove at tail function is invalid")
	}

	//case 2: list is not nil
	res2 := New()
	res2.InsertAtHead(value)
	res2.InsertAtHead(value + 1)
	res2.InsertAtHead(value + 2)
	err = res2.RemoveAtTail()
	if err != nil {
		t.Errorf("%v", err)
	}
	if value == res2.Head.Next.Val {
		t.Errorf("Remove at tail function is invalid")
	}
}

func TestLinkedList_RemoveByIndex(t *testing.T) {
	value := rand.Int()
	res := New()

	//case 1: linkedList is nil
	err := res.RemoveByIndex(0)
	if err == nil {
		t.Errorf("When the linkedList is nil, remove by index function is invalid")
	}

	res.InsertAtTail(value)
	res.InsertAtTail(value + 1)
	res.InsertAtTail(value + 2)

	//case 2: index < 0
	err = res.RemoveByIndex(-1)
	if err == nil {
		t.Errorf("When the index < 0, remove by index function is invalid")
	}

	//case 3: 0 < index < end
	err = res.RemoveByIndex(1)
	if err != nil {
		t.Errorf("%v", err)
	}
	if value+2 != res.Head.Next.Val {
		t.Errorf("When the 0 < index < end, remove by index function is invalid")
	}

	//case 4: end < index
	err = res.RemoveByIndex(5)
	if err == nil {
		t.Errorf("When the end < index, remove by index function is invalid")
	}
}

func TestLinkedList_RemoveByValue(t *testing.T) {
	value := rand.Int()
	res := New()

	//case 1: equals func is nil
	err := res.RemoveByValue(value, nil)
	if err == nil {
		t.Errorf("When the equals function is nil, remove by value function is invalid")
	}

	//case 2: linkedList is nil
	err = res.RemoveByValue(value, equals)
	if err == nil {
		t.Errorf("When the linkedList is nil, remove by value function is invalid")
	}

	//case 3: no error
	res.InsertAtTail(value)
	res.InsertAtTail(value + 1)
	res.InsertAtTail(value + 2)
	err = res.RemoveByValue(value, equals)
	if err != nil {
		t.Errorf("%v", err)
	}
	if value == res.Head.Val {
		t.Errorf("Remove by value function is invalid")
	}
}

func equals(a, b interface{}) bool {
	return a.(int) == b.(int)
}

func TestLinkedList_MoveToHead(t *testing.T) {
	value := rand.Int()
	res := New()

	//case 1: linkedList is nil
	err := res.MoveToHead(1)
	if err == nil {
		t.Errorf("When the linkedList is nil, move to head function is invalid")
	}

	res.InsertAtTail(value)
	res.InsertAtTail(value + 1)
	res.InsertAtTail(value + 2)

	//case 2: index < 0
	err = res.MoveToHead(-1)
	if err == nil {
		t.Errorf("When the index < 0, move to head function is invalid")
	}

	//case 3: 0 < index < end
	err = res.MoveToHead(1)
	if err != nil {
		t.Errorf("When the 0 < index < end, move to head function is invalid")
	}
	if value+1 != res.Head.Val {
		t.Errorf("When the 0 < index < end, move to head function is invalid")
	}

	//case 4: end < index
	err = res.MoveToHead(7)
	if err == nil {
		t.Errorf("When the end < index, move to head function is invalid")
	}
}

func TestLinkedList_MoveToTail(t *testing.T) {
	value := rand.Int()
	res := New()

	//case 1: linkedList is nil
	err := res.MoveToTail(1)
	if err == nil {
		t.Errorf("When the linkedList is nil, move to tail function is invalid")
	}

	res.InsertAtTail(value)
	res.InsertAtTail(value + 1)
	res.InsertAtTail(value + 2)

	//case 2: index < 0
	err = res.MoveToTail(-1)
	if err == nil {
		t.Errorf("When the index < 0, move to tail function is invalid")
	}

	//case 3: 0 < index < end
	err = res.MoveToTail(1)
	if err != nil {
		t.Errorf("When the 0 < index < end, move to tail function is invalid")
	}
	if value+1 != res.Head.Next.Next.Val {
		t.Errorf("When the 0 < index < end, move to tail function is invalid")
	}

	//case 4: end < index
	err = res.MoveToTail(7)
	if err == nil {
		t.Errorf("When the end < index, move to tail function is invalid")
	}
}

func TestLinkedList_SearchByValue(t *testing.T) {
	value := rand.Int()
	res := New()

	//case 1: equals func is nil
	_, err := res.SearchByValue(value, nil)
	if err == nil {
		t.Errorf("When the equals function is nil, search by value function is invalid")
	}

	//case 2: linkedList is nil
	_, err = res.SearchByValue(value, equals)
	if err == nil {
		t.Errorf("When the linkedList is nil, search by value function is invalid")
	}

	res.InsertAtTail(value)
	res.InsertAtTail(value + 1)
	res.InsertAtTail(value + 2)

	//case 3: true
	flag, err := res.SearchByValue(value, equals)
	if err != nil {
		t.Errorf("%v", err)
	}
	if !flag {
		t.Errorf("When the value is exsit, search by value function is invalid")
	}

	//case 4: false
	flag, err = res.SearchByValue(value+3, equals)
	if err != nil {
		t.Errorf("%v", err)
	}
	if flag {
		t.Errorf("When the value is not exsit, search by value function is invalid")
	}
}

func TestLinkedList_SearchByIndex(t *testing.T) {
	res := New()
	value := rand.Int()

	//case 1: linked list is nil
	_, err := res.SearchByIndex(0)
	if err == nil {
		t.Errorf("When the linked list is nil, search by index function is invalid")
	}

	//case 2: index < 0
	res.InsertAtTail(value)
	res.InsertAtTail(value + 1)
	res.InsertAtTail(value + 2)
	_, err = res.SearchByIndex(-1)
	if err == nil {
		t.Errorf("When the index < 0, search by index function is invalid")
	}

	//case 3: 0 < index < end
	node, err := res.SearchByIndex(1)
	if err != nil {
		t.Errorf("When the 0 < index < end, search by index function is invalid")
	}
	if node == nil {
		t.Errorf("When the 0 < index < end, search by index function is invalid")
	}

	//case 4: end < index
	_, err = res.SearchByIndex(5)
	if err == nil {
		t.Errorf("When the end < index, search by index function is invalid")
	}

}

func TestLinkedList_DetectCycle(t *testing.T) {
	res := New()
	value := rand.Int()

	//case 1: list is nil
	flag := res.DetectCycle()
	if flag {
		t.Errorf("When the linked list is nil, detect cycle function is invalid")
	}

	//case 2: false
	res.InsertAtTail(value)
	res.InsertAtTail(value + 1)
	res.InsertAtTail(value + 2)
	flag = res.DetectCycle()
	if flag {
		t.Errorf("When the linked list does not have circle, detect cycle function is invalid")
	}

	//case 3: true
	res.Head.Next.Next.Next = res.Head.Next
	flag = res.DetectCycle()
	if !flag {
		t.Errorf("When the linked list has circle, detect cycle function is invalid")
	}

}

func TestLinkedList_DetectIntersection(t *testing.T) {
	value := rand.Int()
	res1 := New()
	res2 := New()
	res1.InsertAtTail(value)
	res1.InsertAtTail(value + 1)
	res2.InsertAtTail(value)
	res2.InsertAtTail(value + 1)
	//case 1: false
	node := res1.DetectIntersection(res2)
	if node != nil {
		t.Errorf("When the linked lists do not have intersection, detect intersection function is invalid")
	}

	//case 2: true
	newNode := &Node{Val: value, Next: nil}
	res1.Head.Next.Next = newNode
	res2.Head.Next.Next = newNode
	node = res1.DetectIntersection(res2)
	if node == nil {
		t.Errorf("When the linked lists have intersection, detect intersection function is invalid")
	}
}

func TestLinkedList_DetectMiddle(t *testing.T) {
	value := rand.Int()
	res := New()

	//case 1: linked list is nil
	node := res.DetectMiddle()
	if node != nil {
		t.Errorf("When the linked list is nil, detect middle function is invalid")
	}

	//case 2: middle node exist
	res.InsertAtTail(value)
	res.InsertAtTail(value + 1)
	res.InsertAtTail(value + 2)
	node = res.DetectMiddle()
	if node == nil {
		t.Errorf("When the linked list is not nil, detect middle function is invalid")
	}
}

func TestLinkedList_Size(t *testing.T) {
	value := rand.Int()
	res := New()
	res.InsertAtTail(value)
	res.InsertAtTail(value + 1)
	res.InsertAtTail(value + 2)
	size := res.Size()
	if size != 3 {
		t.Errorf("Size function is invalid")
	}
}

func TestLinkedList_Clone(t *testing.T) {
	value := rand.Int()
	res1 := New()
	//case 1: linked list is nil
	res2 := res1.Clone()
	if res2 != nil {
		t.Errorf("When the linked list is nil, clone function is invalid")
	}

	//case 2: linked list is not nil
	res1.InsertAtTail(value)
	res1.InsertAtTail(value + 1)
	res1.InsertAtTail(value + 2)
	res2 = res1.Clone()
	cur1 := res1.Head
	cur2 := res2.Head
	for cur1 != nil && cur2 != nil {
		if cur1.Val != cur2.Val {
			t.Errorf("When the linked list is not nil, clone function is invalid")
		}
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
}

func TestLinkedList_Reverse(t *testing.T) {
	value := rand.Int()
	res := New()
	//case 1: linked list is nil
	err := res.Reverse()
	if err == nil {
		t.Errorf("When the linked list is nil, reverse function is invalid")
	}

	//case 2: linked list is not nil
	res.InsertAtTail(value)
	res.InsertAtTail(value + 1)
	res.InsertAtTail(value + 2)
	err = res.Reverse()
	if err != nil {
		t.Errorf("When the linked list is not nil, reverse function is invalid")
	}
	if res.Head.Val != value+2 || res.Head.Next.Val != value+1 || res.Head.Next.Next.Val != value {
		t.Errorf("When the linked list is not nil, reverse function is invalid")
	}
}

func TestLinkedList_Split(t *testing.T) {
	value := rand.Int()
	res := New()
	//case 1: list is nil
	res1, res2 := res.Split()
	if res1 == nil || res2 != nil {
		t.Errorf("When the linked list is nil, split function is invalid")
	}

	//case 2: list is not nil
	res.InsertAtTail(value)
	res.InsertAtTail(value + 1)
	res.InsertAtTail(value + 2)
	res.InsertAtTail(value + 3)
	res1, res2 = res.Split()
	if res1.Head.Val != value || res1.Head.Next.Val != value+1 || res2.Head.Val != value+2 || res2.Head.Next.Val != value+3 {
		t.Errorf("When the linked list is not nil, split function is invalid")
	}

}
