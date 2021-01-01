package dijkstra

type Graph struct {
	Vertices []Vertex
}

func NewGraph() *Graph {
	return &Graph{
		Vertices: []Vertex{},
	}
}

func (g *Graph) AddNewVertex(ID string) *Vertex {
	for _, ver := range g.Vertices {
		if ver.ID == ID {
			// mean this vertex already exist.
			return &ver
		}
	}
	v := NewVertex(ID)
	g.Vertices = append(g.Vertices, v)
	return &v
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

func (g *Graph) GetVertex(ID string) *Vertex {
	for _, ver := range g.Vertices {
		if ver.ID == ID {
			return &ver
		}
	}
	return nil
}
