package dijkstra

import (
	"testing"
)

func TestDijkstraInitTable(t *testing.T) {
	g := NewGraph()

	g.AddNewVertex("a")
	g.AddNewVertex("b")
	g.AddNewVertex("c")
	g.AddNewVertex("d")
	g.AddNewVertex("e")
	g.AddNewVertex("f")

	g.AppendPath("a", "b", 2)
	g.AppendPath("a", "c", 12)
	g.AppendPath("b", "c", 6)
	g.AppendPath("b", "d", 8)
	g.AppendPath("c", "e", 10)
	g.AppendPath("c", "f", 17)
	g.AppendPath("d", "e", 7)
	g.AppendPath("e", "f", 5)

	g.DrawTableWithPath()
	g.DrawTableWithDistance()
}
