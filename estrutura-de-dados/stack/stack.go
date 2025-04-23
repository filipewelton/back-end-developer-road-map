package stack

type Stack[T any] struct {
	size     int
	capacity int
	data     []any
}

func (q *Stack[T]) New(capacity int) {
	q.capacity = capacity

	q.data = make([]any, capacity)

	for index := range capacity {
		q.data[index] = nil
	}
}

func (s *Stack[T]) Push(element T) {
	if s.capacity == 0 {
		panic("Uninitialized stack")
	} else if s.size == s.capacity {
		panic("The stack is full")
	}

	s.data[s.size] = element
	s.size = s.size + 1
}

func (s *Stack[T]) Pop() T {
	if s.capacity == 0 {
		panic("Uninitialized stack")
	} else if s.size == 0 {
		panic("The stack is empty")
	}

	index := s.size - 1
	removed := s.data[index]
	s.data = s.data[0:index]

	return removed.(T)
}

func (s *Stack[T]) Top() T {
	if s.size == 0 {
		panic("The stack is empty")
	}

	return s.data[s.size-1].(T)
}

func (q *Stack[T]) IsEmpty() bool {
	if q.capacity == 0 {
		panic("Uninitialized stack")
	}

	return q.size == 0
}

func (q *Stack[T]) IsFull() bool {
	if q.capacity == 0 {
		panic("Uninitialized stack")
	}

	return q.capacity == q.size
}
