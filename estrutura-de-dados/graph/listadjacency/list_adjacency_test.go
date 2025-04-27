package listadjacency

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	var graph GraphAdjacencyList[int]

	head := graph.New(1)

	expected := GraphAdjacencyList[int]{
		value: 1,
		next:  nil,
	}

	require.Equal(t, expected, *head)
}

func TestAdd(t *testing.T) {
	var graph GraphAdjacencyList[int]

	head := graph.New(1)
	head.Add(head, 2)

	expected := GraphAdjacencyList[int]{
		value: 1,
		next: &GraphAdjacencyList[int]{
			value: 2,
		},
	}

	require.Equal(t, expected, *head)
}

func TestRemove(t *testing.T) {
	var graph GraphAdjacencyList[int]

	v1 := graph.New(1)
	v2 := v1.Add(v1, 2)
	v2.Add(v2, 3)

	expected := GraphAdjacencyList[int]{
		value: 2,
	}

	v2.Remove(3)
	require.Equal(t, expected, *v2)
}

func TestSearch(t *testing.T) {
	var graph GraphAdjacencyList[int]

	v1 := graph.New(1)
	v2 := v1.Add(v1, 2)
	v3 := v2.Add(v2, 3)

	v3.Add(v3, 4)

	searchResult := v3.Search(4)
	expected := GraphAdjacencyList[int]{value: 4}

	require.Equal(t, expected, *searchResult)
}
