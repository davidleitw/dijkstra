package dijkstra

type Graph struct {
	Vertices []Vertex
}

func NewGraph() *Graph {
	return &Graph{
		Vertices: []Vertex{},
	}
}

func (g *Graph) AddNewVertex() {

}

func (g *Graph) AddVertex(v Vertex) *Vertex {
	for _, ver := range g.Vertices {
		if ver.ID == v.ID {
			// mean this vertex already exist.
			return &ver
		}
	}
	g.Vertices = append(g.Vertices, v)
	return &v
}
