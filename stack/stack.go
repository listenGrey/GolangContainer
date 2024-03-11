package stack

type Stack struct {
	Items []interface{}
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Push(item interface{}) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() interface{} {
	if len(s.Items) == 0 {
		return nil
	}

	item := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return item
}

func (s *Stack) Peek() interface{} {
	if len(s.Items) == 0 {
		return nil
	}
	return s.Items[len(s.Items)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.Items) == 0
}

func (s *Stack) Size() int {
	return len(s.Items)
}
