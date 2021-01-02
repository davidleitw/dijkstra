package dijkstra

import (
	"errors"
)

type Graph struct {
	Vertices []*Vertex
	Vns      []int // vertex number array
	table    map[int]map[int]int64
}

func NewGraph() *Graph {
	return &Graph{
		Vertices: []*Vertex{},
		Vns:      []int{},
		table:    map[int]map[int]int64{},
	}
}

func (g *Graph) AddNewVertex(ID string) *Vertex {
	for _, ver := range g.Vertices {
		if ver.ID == ID {
			// mean this vertex already exist.
			return ver
		}
	}

	v, err := NewVertex(SetID(ID), SetNumber(len(g.Vertices)))
	if err != nil {
		return nil
	}

	g.Vertices = append(g.Vertices, v)
	g.Vns = append(g.Vns, v.number)
	return v
}

func (g *Graph) AddVertex(v *Vertex) *Vertex {
	for _, ver := range g.Vertices {
		if ver.ID == v.ID {
			// mean this vertex already exist.
			return ver
		}
	}
	g.Vertices = append(g.Vertices, v)
	g.Vns = append(g.Vns, v.number)
	return v
}

func (g *Graph) GetVertex(ID string) *Vertex {
	for _, ver := range g.Vertices {
		if ver.ID == ID {
			return ver
		}
	}
	return nil
}

func (g *Graph) getVertexNumber(ID string) int {
	for _, ver := range g.Vertices {
		if ver.ID == ID {
			return ver.number
		}
	}
	return -1
}

func (g *Graph) AppendPath(src, dest string, distance int64) error {
	n1 := g.GetVertex(src)
	n2 := g.GetVertex(dest)
	if n1.number == -1 || n2.number == -1 {
		return errors.New("Please check your src or dest ID.")
	}

	n1.appendPath(n2.number, distance)
	return nil
}
