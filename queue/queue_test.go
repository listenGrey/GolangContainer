package queue

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	mock := &Queue{}
	q := New()
	if reflect.TypeOf(mock) != reflect.TypeOf(q) {
		t.Errorf("New function is invalid")
	}
}

func TestQueue_Push(t *testing.T) {
	value := rand.Int()
	q := New()
	q.Push(value)
	if q.Items[0] != value {
		t.Errorf("Push function is invalid")
	}
}

func TestQueue_Pop(t *testing.T) {

}

func TestQueue_Front(t *testing.T) {

}

func TestQueue_Back(t *testing.T) {

}

func TestQueue_IsEmpty(t *testing.T) {

}

func TestQueue_Size(t *testing.T) {

}
