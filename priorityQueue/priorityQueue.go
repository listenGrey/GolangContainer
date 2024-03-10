package priorityQueue

import "container/heap"

type PriQueue []int

func (q PriQueue) Len() int {
	return len(q)
}

func (q PriQueue) Less(i, j int) bool {
	return q[i] < q[j]
}

func (q PriQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *PriQueue) Push(x interface{}) {
	*q = append(*q, x.(int))
	heap.Init(q)
}

func (q *PriQueue) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[:n-1]
	heap.Init(q)
	return x
}

func (q *PriQueue) Peek() interface{} {
	queue := *q
	n := len(queue)
	return queue[n-1]
}

func (q PriQueue) Init() {
	heap.Init(&q)
}
