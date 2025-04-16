package estruturadedados

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsert(t *testing.T) {
	t.Run("With two elements", func(t *testing.T) {
		var (
			head = LinkedList[int]{
				Value: 1,
				Prev:  nil,
				Next:  nil,
			}
		)

		head.Insert(2)
		require.Nil(t, head.Prev)
		require.NotNil(t, head.Next)
		require.Equal(t, 2, head.Next.Value)
	})

	t.Run("With four elements", func(t *testing.T) {
		var (
			head = LinkedList[int]{
				Value: 1,
				Prev:  nil,
				Next: &LinkedList[int]{
					Value: 2,
					Next: &LinkedList[int]{
						Value: 3,
					},
				},
			}
		)

		head.Insert(4)
		require.Nil(t, head.Prev)
		require.NotNil(t, head.Next)
		require.Equal(t, 4, head.Next.Next.Next.Value)
	})
}

func TestSearch(t *testing.T) {
	t.Run("With valid element", func(t *testing.T) {
		var (
			head = LinkedList[int]{
				Value: 1,
				Prev:  nil,
				Next: &LinkedList[int]{
					Value: 2,
					Next: &LinkedList[int]{
						Value: 3,
					},
				},
			}
		)

		ok, _ := head.Search(3)
		require.True(t, ok)
	})

	t.Run("With invalid element", func(t *testing.T) {
		var (
			head = LinkedList[int]{
				Value: 1,
				Prev:  nil,
				Next: &LinkedList[int]{
					Value: 2,
					Next: &LinkedList[int]{
						Value: 3,
					},
				},
			}
		)

		ok, _ := head.Search(4)
		require.False(t, ok)
	})
}

func TestExclude(t *testing.T) {
	t.Run("With valid element", func(t *testing.T) {
		var (
			head = LinkedList[int]{
				Value: 1,
				Prev:  nil,
				Next: &LinkedList[int]{
					Value: 2,
					Next: &LinkedList[int]{
						Value: 3,
					},
				},
			}
		)

		ok := head.Exclude(2)
		require.True(t, ok)
		require.Equal(t, 3, head.Next.Value)
		require.Nil(t, head.Next.Next)
	})

	t.Run("With invalid element", func(t *testing.T) {
		var head = LinkedList[int]{Value: 1}

		ok := head.Exclude(2)
		require.False(t, ok)
	})
}
