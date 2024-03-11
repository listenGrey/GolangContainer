package stack

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	mock := &Stack{}
	res := New()
	if reflect.TypeOf(mock) != reflect.TypeOf(res) {
		t.Errorf("Creat stack function is invalid")
	}
}

func TestStack_Push(t *testing.T) {
	value := rand.Int()
	mock := &Stack{}
	mock.Items = append(mock.Items, value)
	res := New()
	res.Push(value)
	if mock.Items[0] != res.Items[0] {
		t.Errorf("Push function is invalid")
	}
}

func TestStack_Pop(t *testing.T) {
	value := rand.Int()
	res := New()
	res.Push(value)
	if value != res.Pop() {
		t.Errorf("Pop function is invalid")
	}
}

func TestStack_Peek(t *testing.T) {
	value := rand.Int()
	res := New()
	res.Push(value)
	if value != res.Peek() {
		t.Errorf("Peek function is invalid")
	}
}

func TestStack_IsEmpty(t *testing.T) {
	stack1 := New()
	stack2 := New()
	value := rand.Int()
	stack2.Push(value)

	if !stack1.IsEmpty() {
		t.Errorf("IsEmpty function is invalid")
	}
	if stack2.IsEmpty() {
		t.Errorf("IsEmpty function is invalid")
	}
}

func TestStack_Size(t *testing.T) {
	stack1 := New()
	stack2 := New()
	length := 10
	for i := 0; i < length; i++ {
		stack2.Push(rand.Int())
	}

	if stack1.Size() != 0 {
		t.Errorf("Size function is invalid")
	}
	if stack2.Size() != length {
		t.Errorf("Size function is invalid")
	}
}
