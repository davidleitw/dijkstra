package dijkstra

import (
	"errors"
)

func (g *Graph) setup() error {
	if len(g.Vertices) == 0 {
		return errors.New("please create vertex first.")
	}

	table := make(map[int]map[int]int64)
	for _, vertex := range g.Vertices {
		if len(vertex.weights) == 0 {
			continue
		}

		t := make(map[int]int64)
		for targetNumber, distance := range vertex.weights {
			t[targetNumber] = distance
		}
		table[vertex.number] = t
	}

	g.table = table
	return nil
}

func (g *Graph) DrawTable() error {
	err := g.setup()
	if err != nil {
		return err
	}

	num := len(g.Vns)
	for _, vertex := range g.Vertices {
		src := vertex.number

		for {
			idx := (src + 1) % num
			if g.Vns[idx] == src {
				break
			}

		}
	}

	return nil
}
