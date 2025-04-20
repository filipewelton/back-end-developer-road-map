package estruturadedados

import "slices"

type Array[T any] struct {
	size        int
	currentSize int
	data        []any
}

func (a *Array[T]) New(size int) {
	a.size = size
	a.currentSize = 0
	a.data = make([]any, size)

	for index := range a.data {
		a.data[index] = nil
	}
}

func (a *Array[T]) Insert(element T) {
	if a.size == 0 {
		panic("Uninitialized array")
	} else if a.currentSize == a.size {
		panic("Array full")
	}

	a.data[a.currentSize] = element
	a.currentSize = a.currentSize + 1
}

func (a *Array[T]) Exclude(element T) bool {
	if a.currentSize == 0 {
		panic("Array is empty")
	}

	for index, elem := range a.data {
		if any(elem) == any(element) {
			a.data = slices.Delete(a.data, index, index+1)
			a.currentSize = a.currentSize - 1

			return true
		}
	}

	return false
}

func (a *Array[T]) Search(element T) int {
	for index, elem := range a.data {
		if any(elem) == any(element) {
			return index
		}
	}

	return -1
}
