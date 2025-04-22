package queue

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
		Queue[int]{
			size:     0,
			capacity: 1,
			data:     make([]any, 1),
		},
		Queue[int]{
			size:     0,
			capacity: 2,
			data:     make([]any, 2),
		},
	}

	for i, testName := range testNames {
		t.Run(testName, func(t *testing.T) {
			var queue Queue[int]

			queue.New(capacities[i])
			require.Equal(t, expected[i], queue)
		})
	}
}

func TestEnqueue(t *testing.T) {
	t.Run("when not initialized", func(t *testing.T) {
		var queue Queue[int]

		expectedError := "Uninitialized queue"

		require.PanicsWithValue(t, expectedError, func() {
			queue.Enqueue(1)
		})
	})

	t.Run("when the queue is full", func(t *testing.T) {
		var queue Queue[int]

		queue.New(1)
		queue.Enqueue(1)

		expectedError := "The queue is full"

		require.PanicsWithValue(t, expectedError, func() {
			queue.Enqueue(2)
		})
	})

	t.Run("when queued", func(t *testing.T) {
		var queue Queue[int]

		queue.New(1)
		queue.Enqueue(1)
		require.Equal(t, queue.capacity, queue.size)
	})
}

func TestDequeue(t *testing.T) {
	t.Run("when not initialized", func(t *testing.T) {
		var queue Queue[int]

		expectedError := "Uninitialized queue"

		require.PanicsWithValue(t, expectedError, func() {
			queue.Dequeue()
		})
	})

	t.Run("when the queue is empty", func(t *testing.T) {
		var queue Queue[int]

		queue.New(1)

		expectedError := "The queue is empty"

		require.PanicsWithValue(t, expectedError, func() {
			queue.Dequeue()
		})
	})

	t.Run("when removed from the queue", func(t *testing.T) {
		var queue Queue[int]

		queue.New(2)
		queue.Enqueue(1)
		queue.Enqueue(2)

		removed := queue.Dequeue()

		require.Equal(t, 1, removed)
	})
}

func TestPeek(t *testing.T) {
	t.Run("when the queue is empty", func(t *testing.T) {
		var queue Queue[int]

		expectedError := "The queue is empty"

		require.PanicsWithValue(t, expectedError, func() {
			queue.Peek()
		})
	})

	t.Run("when spied", func(t *testing.T) {
		var queue Queue[int]

		queue.New(1)
		queue.Enqueue(1)

		element := queue.Peek()

		require.Equal(t, 1, element)
	})
}

func TestSize(t *testing.T) {
	t.Run("when return size", func(t *testing.T) {
		var queue Queue[int]

		queue.New(2)
		queue.Enqueue(1)

		size := queue.Size()

		require.Equal(t, 1, size)
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("when the queue is empty", func(t *testing.T) {
		var queue Queue[int]

		queue.New(1)

		result := queue.IsEmpty()

		require.True(t, result)
	})

	t.Run("when the queue is not empty", func(t *testing.T) {
		var queue Queue[int]

		queue.New(1)
		queue.Enqueue(1)

		result := queue.IsEmpty()

		require.False(t, result)
	})
}

func TestIsFull(t *testing.T) {
	t.Run("when the queue is empty", func(t *testing.T) {
		var queue Queue[int]

		queue.New(1)
		result := queue.IsFull()

		require.False(t, result)
	})

	t.Run("when the queue is full", func(t *testing.T) {
		var queue Queue[int]

		queue.New(1)
		queue.Enqueue(1)

		result := queue.IsFull()

		require.True(t, result)
	})
}
