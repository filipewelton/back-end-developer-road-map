// Implementação de um grafo utilizando matriz de adjacência.

package adjacencymatrix

type GraphAdjacencyMatrix struct {
	data             [][]int
	numberOfVertices int
}

func (g *GraphAdjacencyMatrix) New() {
	g.data = make([][]int, 1)
	g.data[0] = []int{0}
	g.numberOfVertices = 1
}

func (g *GraphAdjacencyMatrix) Add(parentVertex int, newVertex int) {
	if parentVertex > g.numberOfVertices {
		panic("Invalid parent vertex")
	} else if newVertex > g.numberOfVertices+1 {
		panic("Invalid new vertex")
	}

	g.data = append(g.data, make([]int, g.numberOfVertices))
	g.numberOfVertices++

	for i := range g.numberOfVertices {
		g.data[i] = append(g.data[i], 0)

		if i == parentVertex {
			g.data[i][newVertex] = 1
		} else if i == newVertex {
			g.data[i][parentVertex] = 1
		}
	}
}

func (g *GraphAdjacencyMatrix) Remove(vertex int) {
	if vertex > g.numberOfVertices {
		panic("Invalid vertex")
	}

	for i, edge := range g.data[vertex] {
		if edge == 1 {
			g.data[i][edge] = 0
		}
	}

	j := g.numberOfVertices - vertex
	updatedData := make([][]int, g.numberOfVertices-j)

	for i := range g.numberOfVertices {
		if i >= vertex {
			continue
		}

		matrix := g.data[i][0 : len(g.data[i])-j]
		updatedData[i] = matrix
	}

	g.data = updatedData
}

func (g *GraphAdjacencyMatrix) Search(vertex int) []int {
	return g.data[vertex]
}
