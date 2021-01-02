package dijkstra

import (
	"fmt"
	"testing"
)

func TestDijkstraInitTable(t *testing.T) {
	g := NewGraph()

	v0, _ := NewVertex(SetID("v0"), SetNumber(0))
	v1, _ := NewVertex(SetID("v1"), SetNumber(1))
	v2, _ := NewVertex(SetID("v2"), SetNumber(2))

	g.AddVertex(v0)
	g.AddVertex(v1)
	g.AddVertex(v2)

	_ = g.AppendPath("v0", "v1", 20)
	_ = g.AppendPath("v1", "v0", 40)
	_ = g.AppendPath("v2", "v0", 55)
	_ = g.AppendPath("v2", "v1", 72)

	err := g.DrawTable()
	if err != nil {
		t.Error(err)
	}

	for src, row := range g.table {
		fmt.Printf("src = %d", src)

		for target, distance := range row {
			fmt.Printf(" => (%d, %d) ", target, distance)
		}
		fmt.Println("")
	}
}
