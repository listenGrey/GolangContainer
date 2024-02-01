package queue

type Queue struct {
	Items []interface{}
}

func New() *Queue {
	return &Queue{}
}

func (q *Queue) Push(item interface{}) {
	q.Items = append(q.Items, item)
}

func (q *Queue) Pop() {
	if len(q.Items) > 0 {
		q.Items = q.Items[1:]
	}
}

func (q *Queue) Front() interface{} {
	if len(q.Items) == 0 {
		return nil
	}
	return q.Items[0]
}

func (q Queue) Back() interface{} {
	if len(q.Items) == 0 {
		return nil
	}
	return q.Items[len(q.Items)-1]
}

func (q *Queue) IsEmpty() bool {
	return len(q.Items) == 0
}

func (q *Queue) Size() int {
	return len(q.Items)
}
