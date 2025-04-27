package adjacencymatrix

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	var graph GraphAdjacencyMatrix

	graph.New()
	require.IsType(t, [][]int{}, graph.data)
}

func TestAdd(t *testing.T) {
	var graph GraphAdjacencyMatrix

	graph.New()
	graph.Add(0, 1)
	graph.Add(1, 2)
	require.Equal(t, 1, graph.data[0][1])
	require.Equal(t, 1, graph.data[1][0])
}

func TestRemove(t *testing.T) {
	var graph GraphAdjacencyMatrix

	graph.New()
	graph.Add(0, 1)
	graph.Add(1, 2)
	graph.Add(2, 3)
	graph.Remove(2)
	require.Equal(t, 1, graph.data[0][1])
	require.Equal(t, 1, graph.data[1][0])
}

func TestSearch(t *testing.T) {
	var graph GraphAdjacencyMatrix

	graph.New()
	graph.Add(0, 1)
	graph.Add(1, 2)
	graph.Add(2, 3)

	result := graph.Search(2)
	expected := []int{0, 1, 0, 1}

	require.Equal(t, expected, result)
}
