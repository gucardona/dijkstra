package graph

type Edge struct {
	To     string
	Weight int
}

type Graph struct {
	Vertices map[string][]Edge
}

func NewGraph() *Graph {
	return &Graph{
		Vertices: make(map[string][]Edge),
	}
}

func (g *Graph) AddEdge(from, to string, weight int) {
	g.Vertices[from] = append(g.Vertices[from], Edge{To: to, Weight: weight})
	g.Vertices[to] = append(g.Vertices[to], Edge{To: from, Weight: weight})
}

func (g *Graph) GetNeighbors(vertex string) []Edge {
	return g.Vertices[vertex]
}

func (g *Graph) GetAllVertices() []string {
	vertices := make([]string, 0, len(g.Vertices))
	for v := range g.Vertices {
		vertices = append(vertices, v)
	}
	return vertices
}
