// Implementação de um grafo utilizando lista de adjacência.

package listadjacency

type GraphAdjacencyList[T any] struct {
	value T
	next  *GraphAdjacencyList[T]
}

func (g *GraphAdjacencyList[T]) New(value T) *GraphAdjacencyList[T] {
	return &GraphAdjacencyList[T]{
		value: value,
	}
}

func (g *GraphAdjacencyList[T]) Add(
	parentVertex *GraphAdjacencyList[T],
	value T,
) *GraphAdjacencyList[T] {
	vertex := GraphAdjacencyList[T]{
		value: value,
	}

	*parentVertex = GraphAdjacencyList[T]{
		value: parentVertex.value,
		next:  &vertex,
	}

	return &vertex
}

func (g *GraphAdjacencyList[T]) Remove(value T) {
	handleRemoval(g, &value)
}

func handleRemoval[T any](head *GraphAdjacencyList[T], value *T) any {
	if any(head.value) == any(*value) {
		panic("Cannot remove the head of the list")
	} else if head.next == nil {
		panic("Vertex not found")
	} else if any(head.next.value) == any(*value) {
		head.next = nil
		return nil
	}

	return handleRemoval(head.next, value)
}

func (g *GraphAdjacencyList[T]) Search(vertext T) *GraphAdjacencyList[T] {
	if any(g.value) == any(vertext) {
		return g
	} else if g.next == nil {
		return nil
	}

	return g.next.Search(vertext)
}
