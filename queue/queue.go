package queue

type Queue struct {
	items []interface{}
}

func New() *Queue {
	return &Queue{}
}

func (q *Queue) Push(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Pop() {
	if len(q.items) > 0 {
		q.items = q.items[1:]
	}
}

func (q *Queue) Front() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	return q.items[0]
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Size() int {
	return len(q.items)
}
