package stack

import (
	"fmt"

	"testing"
)

func TestNew(t *testing.T) {
	mock := &Stack{}
	res := New()
	if mock != res {
		fmt.Println("Creat stack function is valid")
	}
}

func TestStack_Push(t *testing.T) {
	//value := rand.Int()

}

func TestStack_Pop(t *testing.T) {

}

func TestStack_Peek(t *testing.T) {

}

func TestStack_IsEmpty(t *testing.T) {

}

func TestStack_Size(t *testing.T) {

}
