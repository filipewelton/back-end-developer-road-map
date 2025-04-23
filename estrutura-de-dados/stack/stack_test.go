package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	testNames := []string{
		"should boot with capacity 1",
		"should boot with capacity 2",
	}

	capacities := []int{1, 2}

	expected := []any{
		Stack[int]{
			size:     0,
			capacity: 1,
			data:     make([]any, 1),
		},
		Stack[int]{
			size:     0,
			capacity: 2,
			data:     make([]any, 2),
		},
	}

	for i, testName := range testNames {
		t.Run(testName, func(t *testing.T) {
			var stack Stack[int]

			stack.New(capacities[i])
			require.Equal(t, expected[i], stack)
		})
	}
}

func TestPush(t *testing.T) {
	t.Run("when not initialized", func(t *testing.T) {
		var stack Stack[int]

		expectedError := "Uninitialized stack"

		require.PanicsWithValue(t, expectedError, func() {
			stack.Push(1)
		})
	})

	t.Run("when the stack is full", func(t *testing.T) {
		var stack Stack[int]

		stack.New(1)
		stack.Push(1)

		expectedError := "The stack is full"

		require.PanicsWithValue(t, expectedError, func() {
			stack.Push(2)
		})
	})

	t.Run("when queued", func(t *testing.T) {
		var stack Stack[int]

		stack.New(1)
		stack.Push(1)
		require.Equal(t, stack.capacity, stack.size)
	})
}

func TestPop(t *testing.T) {
	t.Run("when not initialized", func(t *testing.T) {
		var stack Stack[int]

		expectedError := "Uninitialized stack"

		require.PanicsWithValue(t, expectedError, func() {
			stack.Pop()
		})
	})

	t.Run("when the stack is empty", func(t *testing.T) {
		var stack Stack[int]

		stack.New(1)

		expectedError := "The stack is empty"

		require.PanicsWithValue(t, expectedError, func() {
			stack.Pop()
		})
	})

	t.Run("when removed from the stack", func(t *testing.T) {
		var stack Stack[int]

		stack.New(2)
		stack.Push(1)
		stack.Push(2)

		removed := stack.Pop()

		require.Equal(t, 2, removed)
	})
}

func TestTop(t *testing.T) {
	t.Run("when the stack is empty", func(t *testing.T) {
		var stack Stack[int]

		expectedError := "The stack is empty"

		require.PanicsWithValue(t, expectedError, func() {
			stack.Top()
		})
	})

	t.Run("when spied", func(t *testing.T) {
		var stack Stack[int]

		stack.New(1)
		stack.Push(1)

		element := stack.Top()

		require.Equal(t, 1, element)
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("when the stack is empty", func(t *testing.T) {
		var stack Stack[int]

		stack.New(1)

		result := stack.IsEmpty()

		require.True(t, result)
	})

	t.Run("when the stack is not empty", func(t *testing.T) {
		var stack Stack[int]

		stack.New(1)
		stack.Push(1)

		result := stack.IsEmpty()

		require.False(t, result)
	})
}

func TestIsFull(t *testing.T) {
	t.Run("when the stack is empty", func(t *testing.T) {
		var stack Stack[int]

		stack.New(1)
		result := stack.IsFull()

		require.False(t, result)
	})

	t.Run("when the stack is full", func(t *testing.T) {
		var stack Stack[int]

		stack.New(1)
		stack.Push(1)

		result := stack.IsFull()

		require.True(t, result)
	})
}
