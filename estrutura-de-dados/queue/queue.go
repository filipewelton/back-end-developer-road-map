package queue

type Queue[T any] struct {
	size     int
	capacity int
	data     []any
}

func (q *Queue[T]) New(capacity int) {
	q.capacity = capacity

	q.data = make([]any, capacity)

	for index := range capacity {
		q.data[index] = nil
	}
}

func (q *Queue[T]) Enqueue(element T) {
	if q.capacity == 0 {
		panic("Uninitialized queue")
	} else if q.size == q.capacity {
		panic("The queue is full")
	}

	q.data[q.size] = element
	q.size = q.size + 1
}

func (q *Queue[T]) Dequeue() T {
	if q.capacity == 0 {
		panic("Uninitialized queue")
	} else if q.size == 0 {
		panic("The queue is empty")
	}

	removed := q.data[0]
	q.data = q.data[1:]
	q.size = q.size - 1

	return removed.(T)
}

func (q *Queue[T]) Peek() T {
	if q.size == 0 {
		panic("The queue is empty")
	}

	return q.data[0].(T)
}

func (q *Queue[T]) Size() int {
	return q.size
}

func (q *Queue[T]) IsEmpty() bool {
	if q.capacity == 0 {
		panic("Uninitialized queue")
	}

	return q.size == 0
}

func (q *Queue[T]) IsFull() bool {
	if q.capacity == 0 {
		panic("Uninitialized queue")
	}

	return q.capacity == q.size
}
