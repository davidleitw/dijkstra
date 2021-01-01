package dijkstra

import "testing"

func TestGetVertexFromGraph(t *testing.T) {
	graph := NewGraph()

	graph.AddNewVertex("test")
	v1 := graph.GetVertex("test")
	if v1 == nil || v1.ID != "test" {
		t.Error("vertex ID should be test.")
	}

	v2 := graph.GetVertex("Hello")
	if v2 != nil {
		t.Error("vertex is not exist, v2 should be nil.")
	}
}

func TestAddVertex(t *testing.T) {
	graph := NewGraph()

	graph.AddNewVertex("test1")
	v1 := graph.GetVertex("test1")
	if v1 == nil || v1.ID != "test1" {
		t.Error("vertex ID should be test1.")
	}
	v2 := graph.AddNewVertex("test1")
	if v2.ID != v1.ID {
		t.Error("vertex ID error.")
	}

	v3 := NewVertex("test2")
	graph.AddVertex(v3)
	v4 := graph.GetVertex("test2")
	if v4 == nil || v4.ID != v3.ID {
		t.Error("vertex ID should be test2.")
	}

	v5 := graph.AddVertex(v3)
	if v5.ID != v4.ID {
		t.Error("vertex ID error.")
	}
}
