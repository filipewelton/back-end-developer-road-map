package estruturadedados

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	var array Array[int]

	array.New(10)

	require.Equal(t, 10, array.size)
	require.Equal(t, 0, array.currentSize)
	require.IsType(t, []any{}, array.data)
}

func TestInsert(t *testing.T) {
	t.Run("With free space", func(t *testing.T) {
		var array Array[int]

		array.New(5)
		array.Insert(1)
		require.Equal(t, 5, array.size)
		require.Equal(t, 1, array.currentSize)
	})

	t.Run("With full array", func(t *testing.T) {
		var array Array[int]

		array.New(1)
		array.Insert(1)

		require.Panics(t, func() {
			array.Insert(2)
		})
	})
}

func TestExclude(t *testing.T) {
	t.Run("With existent element", func(t *testing.T) {
		var array Array[int]

		array.New(1)
		array.Insert(1)

		result := array.Exclude(1)

		require.True(t, result)
	})

	t.Run("With non-existent element", func(t *testing.T) {
		var array Array[int]

		array.New(1)
		array.Insert(0)

		result := array.Exclude(1)

		require.False(t, result)
	})
}

func TestSearch(t *testing.T) {
	t.Run("With existent element", func(t *testing.T) {
		var array Array[int]

		array.New(2)
		array.Insert(1)
		array.Insert(2)

		result := array.Search(2)

		require.Equal(t, 1, result)
	})

	t.Run("With non-existent element", func(t *testing.T) {
		var array Array[int]

		array.New(2)
		array.Insert(1)

		result := array.Search(2)

		require.Equal(t, -1, result)
	})
}
