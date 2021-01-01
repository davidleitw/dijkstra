package dijkstra

type Vertex struct {
	// vertex ID for single node.
	ID      string
	weights map[string]int64
}

func NewVertex(ID string) Vertex {
	return Vertex{
		ID:      ID,
		weights: map[string]int64{},
	}
}

func (v *Vertex) SetWeights(weights map[string]int64) {
	v.weights = weights
}

func (v *Vertex) AppendPath(name string, distance int64) {
	v.weights[name] = distance
}

func (v *Vertex) DeletePath(name string, distance int64) {
	delete(v.weights, name)
}
