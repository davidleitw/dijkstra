package dijkstra

import "errors"

type Graph struct {
	Vertices        []*Vertex
	table           [][]info
	conversionTable map[int]string
}

type info struct {
	distance int64
	path     string
}

func NewGraph() *Graph {
	return &Graph{
		Vertices:        []*Vertex{},
		table:           [][]info{},
		conversionTable: map[int]string{},
	}
}

func (g *Graph) AddNewVertex(ID string) *Vertex {
	for _, vertex := range g.Vertices {
		if vertex.ID == ID {
			// mean this vertex already exist.
			return vertex
		}
	}

	v, err := NewVertex(SetID(ID), SetNumber(len(g.Vertices)))
	if err != nil {
		return nil
	}

	g.Vertices = append(g.Vertices, v)
	g.conversionTable[v.number] = v.ID
	return v
}

func (g *Graph) getVertex(ID string) *Vertex {
	for _, vertex := range g.Vertices {
		if vertex.ID == ID {
			return vertex
		}
	}
	return nil
}

func (g *Graph) AppendPath(src, dest string, distance int64) error {
	if distance <= 0 {
		return errors.New("Distance should > 0.")
	}

	v1 := g.getVertex(src)
	if v1 == nil {
		return errors.New("src ID error, please check your source vertex ID.")
	}

	v2 := g.getVertex(dest)
	if v2 == nil {
		return errors.New("destination ID error, please check your destination ID.")
	}

	v1.appendPath(v2.number, distance)
	return nil
}
