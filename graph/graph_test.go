package graph

import "testing"

func TestGraph_AddVertex(t *testing.T) {
	// Arrange
	graph := &Graph{
		Vertices: make(map[string]*Vertex),
	}
	// Act
	graph.AddVertex("AX1")
	// Assert
	if graph.Count != 1 {
		t.Errorf("Expected 1, but got %d", graph.Count)
	}
	if len(graph.Vertices) != 1 {
		t.Errorf("Expected 1, but got %d", len(graph.Vertices))
	}
}

func TestGraph_AddVertex_Duplicate(t *testing.T) {
	// Arrange
	graph := &Graph{
		Vertices: make(map[string]*Vertex),
	}
	// Act
	graph.AddVertex("AX1")
	graph.AddVertex("AX1")
	// Assert
	if graph.Count != 1 {
		t.Errorf("Expected 1, but got %d", graph.Count)
	}
	if len(graph.Vertices) != 1 {
		t.Errorf("Expected 1, but got %d", len(graph.Vertices))
	}
}

func TestGraph_AddConnection(t *testing.T) {
	// Arrange
	graph := &Graph{
		Vertices: make(map[string]*Vertex),
	}
	graph.AddVertex("AX1")
	graph.AddVertex("AX2")
	// Act
	graph.AddConnection("AX1", "AX2", 3)
	// Assert
	if len(graph.Vertices["AX1"].Connections) != 1 {
		t.Errorf("Expected 1, but got %d", len(graph.Vertices["AX1"].Connections))
	}
	if len(graph.Vertices["AX2"].Connections) != 1 {
		t.Errorf("Expected 1, but got %d", len(graph.Vertices["AX2"].Connections))
	}
}

func TestGraph_AddConnection_Duplicate(t *testing.T) {
	// Arrange
	graph := &Graph{
		Vertices: make(map[string]*Vertex),
	}
	graph.AddVertex("AX1")
	graph.AddVertex("AX2")
	// Act
	graph.AddConnection("AX1", "AX2", 3)
	graph.AddConnection("AX1", "AX2", 3)
	// Assert
	if len(graph.Vertices["AX1"].Connections) != 1 {
		t.Errorf("Expected 1, but got %d", len(graph.Vertices["AX1"].Connections))
	}
	if len(graph.Vertices["AX2"].Connections) != 1 {
		t.Errorf("Expected 1, but got %d", len(graph.Vertices["AX2"].Connections))
	}
}

func TestGraph_AddConnection_AddsWeights(t *testing.T) {
	// Arrange
	graph := &Graph{
		Vertices: make(map[string]*Vertex),
	}
	graph.AddVertex("AX1")
	graph.AddVertex("AX2")
	// Act
	graph.AddConnection("AX1", "AX2", 3)
	// Assert
	if graph.Vertices["AX1"].Connections["AX2"] != 3 {
		t.Errorf("Expected 3, but got %d", graph.Vertices["AX1"].Connections["AX2"])
	}
	if graph.Vertices["AX2"].Connections["AX1"] != 3 {
		t.Errorf("Expected 3, but got %d", graph.Vertices["AX2"].Connections["AX1"])
	}

}

func TestNewGraph_EmptyAdjacencyList(t *testing.T) {
	// Arrange
	adjacencyList := []string{}
	// Act
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	_ = NewGraph(adjacencyList)
}

func TestNewGraph_MakesVertices(t *testing.T) {
	// Arrange
	adjacencyList := []string{
		"AX1,AX2,AX3,AX4,AX5",
		"AX1,AX4:3,AX2:3,AX3:6",
		"AX2,AX1:3,AX3:3,AX4:6",
		"AX3,AX2:3,AX1:6,AX4:4",
		"AX4,AX1:3,AX2:6,AX3:4,AX5:15",
		"AX5,AX4:15",
	}
	// Act
	graph := NewGraph(adjacencyList)
	// Assert
	if graph.Count != 5 {
		t.Errorf("Expected 5, but got %d", graph.Count)
	}
	if len(graph.Vertices) != 5 {
		t.Errorf("Expected 5, but got %d", len(graph.Vertices))
	}
}

func TestNewGraph_MakesConnections(t *testing.T) {
	// Arrange
	adjacencyList := []string{
		"AX1,AX2,AX3,AX4,AX5",
		"AX1,AX4:3,AX2:3,AX3:6",
		"AX2,AX1:3,AX3:3,AX4:6",
		"AX3,AX2:3,AX1:6,AX4:4",
		"AX4,AX1:3,AX2:6,AX3:4,AX5:15",
		"AX5,AX4:15",
	}
	// Act
	graph := NewGraph(adjacencyList)
	// Assert
	if len(graph.Vertices["AX1"].Connections) != 3 {
		t.Errorf("Expected 3, but got %d", len(graph.Vertices["AX1"].Connections))
	}

	if len(graph.Vertices["AX2"].Connections) != 3 {
		t.Errorf("Expected 3, but got %d", len(graph.Vertices["AX2"].Connections))
	}

	if len(graph.Vertices["AX3"].Connections) != 3 {
		t.Errorf("Expected 3, but got %d", len(graph.Vertices["AX3"].Connections))
	}

	if len(graph.Vertices["AX4"].Connections) != 4 {
		t.Errorf("Expected 4, but got %d", len(graph.Vertices["AX4"].Connections))
	}

	if len(graph.Vertices["AX5"].Connections) != 1 {
		t.Errorf("Expected 1, but got %d", len(graph.Vertices["AX5"].Connections))
	}
}

func TestNewGraph_InvalidConnection(t *testing.T) {
	// Arrange
	adjacencyList := []string{
		"AX1,AX2,AX3,AX4,AX5",
		"AX1,AX4:3,AX2:3,AX3:6",
		"AX2,AX1:3,AX3:3,AX4:6",
		"AX3,AX2:3,AX1:6,AX4:4",
		"AX4,AX1:3,AX2:6,AX3:4,AX5:15",
		"AX5,AX4:15",
		"AX6,AX4:15",
	}
	// Act
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	_ = NewGraph(adjacencyList)
}

func TestNewGraph_InvalidWeight(t *testing.T) {
	// Arrange
	adjacencyList := []string{
		"AX1,AX2,AX3,AX4,AX5",
		"AX1,AX4:3,AX2:3,AX3:6",
		"AX2,AX1:3,AX3:3,AX4:6",
		"AX3,AX2:3,AX1:6,AX4:4",
		"AX4,AX1:3,AX2:6,AX3:4,AX5:15",
		"AX5,AX4:15",
		"AX6,AX4:15",
	}
	// Act
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	_ = NewGraph(adjacencyList)
}
