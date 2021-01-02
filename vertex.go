package dijkstra

import (
	"errors"
)

type Vertex struct {
	// vertex ID for single node.
	ID      string
	number  int
	weights map[int]int64
}

type VertexOption interface {
	apply(*Vertex)
}

func NewVertex(opts ...VertexOption) (*Vertex, error) {
	v := &Vertex{
		ID:      "",
		number:  -1,
		weights: map[int]int64{},
	}

	for _, opt := range opts {
		opt.apply(v)
	}

	if v.ID == "" {
		return nil, errors.New("Setting vertex ID is necessarily.")
	}
	if v.number == -1 {
		return nil, errors.New("Setting vertex number is necessarily.")
	}

	return v, nil
}

func (v *Vertex) ApplyOptions(opts ...VertexOption) {
	for _, opt := range opts {
		opt.apply(v)
	}
}

type optFunc func(*Vertex)

func (f optFunc) apply(v *Vertex) {
	f(v)
}

func SetID(ID string) optFunc {
	return optFunc(func(v *Vertex) {
		v.ID = ID
	})
}

func SetNumber(number int) optFunc {
	return optFunc(func(v *Vertex) {
		v.number = number
	})
}

func SetWeights(weights map[int]int64) optFunc {
	return optFunc(func(v *Vertex) {
		v.weights = weights
	})
}

func (v *Vertex) appendPath(dest int, distance int64) {
	v.weights[dest] = distance
}

func (v *Vertex) deletePath(dest int, distance int64) {
	delete(v.weights, dest)
}
