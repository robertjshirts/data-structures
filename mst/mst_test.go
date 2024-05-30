package mst

import (
	"github.com/robertjshirts/data-structures/graph"
	"strings"
	"testing"
)

const inputA = "AX1,AX2,AX3,AX4,AX5\nAX1,AX4:3,AX2:3,AX3:6\nAX2,AX1:3,AX3:3,AX4:6\nAX3,AX2:3,AX1:6,AX4:4\nAX4,AX1:3,AX2:6,AX3:4,AX5:15\nAX5,AX4:15"
const inputB = "AX10,AX11,AX12,AX99,AX100\nAX10,AX11:2,AX12:4\nAX11,AX10:2,AX12:2\nAX12,AX10:4,AX11:2,AX99:4\nAX99,AX12:4,AX100:15\nAX100,AX99:15"

func TestPrim_NilGraph(t *testing.T) {
	// Run Prim's algorithm
	mst := Prim(nil)

	// Check the total weight
	if Weight(mst) != 0 {
		t.Errorf("Expected total weight of 0, got %v", Weight(mst))
	}
}

func TestPrim_EmptyGraph(t *testing.T) {
	// Create a new graph
	inputGraph := graph.EmptyGraph()

	// Run Prim's algorithm
	mst := Prim(inputGraph)

	// Check the total weight
	if Weight(mst) != 0 {
		t.Errorf("Expected total weight of 0, got %v", Weight(mst))
	}
}

func TestPrim_inputA_FindsCorrectWeight(t *testing.T) {
	// Create a new graph
	inputGraph := graph.NewGraph(strings.Split(inputA, "\n"))

	// Run Prim's algorithm
	mst := Prim(inputGraph)

	// Check the total weight
	if Weight(mst) != 24 {
		t.Errorf("Expected total weight of 16, got %v", Weight(mst))
	}
}

func TestPrim_inputA_FindsAllVertices(t *testing.T) {
	// Create a new graph
	inputGraph := graph.NewGraph(strings.Split(inputA, "\n"))

	// Run Prim's algorithm
	mst := Prim(inputGraph)

	// Check the vertices
	if len(mst.Vertices) != 5 {
		t.Errorf("Expected 5 vertices, got %v", len(mst.Vertices))
	}
}

func TestPrim_inputA_FindsCorrectVertices(t *testing.T) {
	// Create a new graph
	inputGraph := graph.NewGraph(strings.Split(inputA, "\n"))

	// Run Prim's algorithm
	mst := Prim(inputGraph)

	// Check the vertices
	if mst.Vertices["AX1"] == nil {
		t.Errorf("Expected AX1 to be in the MST")
	}
	if mst.Vertices["AX2"] == nil {
		t.Errorf("Expected AX2 to be in the MST")
	}
	if mst.Vertices["AX3"] == nil {
		t.Errorf("Expected AX3 to be in the MST")
	}
	if mst.Vertices["AX4"] == nil {
		t.Errorf("Expected AX4 to be in the MST")
	}
	if mst.Vertices["AX5"] == nil {
		t.Errorf("Expected AX5 to be in the MST")
	}
}

func TestPrim_inputA_FindsAllConnections(t *testing.T) {
	// Create a new graph
	inputGraph := graph.NewGraph(strings.Split(inputA, "\n"))

	// Run Prim's algorithm
	mst := Prim(inputGraph)

	// Check the connections
	if len(mst.Vertices["AX1"].Connections) != 2 {
		t.Errorf("Expected 2 connection for AX1, got %v", len(mst.Vertices["AX1"].Connections))
	}
	if len(mst.Vertices["AX2"].Connections) != 2 {
		t.Errorf("Expected 2 connection for AX2, got %v", len(mst.Vertices["AX2"].Connections))
	}
	if len(mst.Vertices["AX3"].Connections) != 1 {
		t.Errorf("Expected 1 connection for AX3, got %v", len(mst.Vertices["AX3"].Connections))
	}
	if len(mst.Vertices["AX4"].Connections) != 2 {
		t.Errorf("Expected 2 connections for AX4, got %v", len(mst.Vertices["AX4"].Connections))
	}
	if len(mst.Vertices["AX5"].Connections) != 1 {
		t.Errorf("Expected 1 connection for AX5, got %v", len(mst.Vertices["AX5"].Connections))
	}
}

func TestPrim_inputB_FindsCorrectWeight(t *testing.T) {
	// Create a new graph
	inputGraph := graph.NewGraph(strings.Split(inputB, "\n"))

	// Run Prim's algorithm
	mst := Prim(inputGraph)

	// Check the total weight
	if Weight(mst) != 23 {
		t.Errorf("Expected total weight of 23, got %v", Weight(mst))
	}
}

func TestPrim_inputB_FindsAllVertices(t *testing.T) {
	// Create a new graph
	inputGraph := graph.NewGraph(strings.Split(inputB, "\n"))

	// Run Prim's algorithm
	mst := Prim(inputGraph)

	// Check the vertices
	if len(mst.Vertices) != 5 {
		t.Errorf("Expected 5 vertices, got %v", len(mst.Vertices))
	}
}

func TestPrim_inputB_FindsCorrectVertices(t *testing.T) {
	// Create a new graph
	inputGraph := graph.NewGraph(strings.Split(inputB, "\n"))

	// Run Prim's algorithm
	mst := Prim(inputGraph)

	// Check the vertices
	if mst.Vertices["AX10"] == nil {
		t.Errorf("Expected AX10 to be in the MST")
	}
	if mst.Vertices["AX11"] == nil {
		t.Errorf("Expected AX11 to be in the MST")
	}
	if mst.Vertices["AX12"] == nil {
		t.Errorf("Expected AX12 to be in the MST")
	}
	if mst.Vertices["AX99"] == nil {
		t.Errorf("Expected AX99 to be in the MST")
	}
	if mst.Vertices["AX100"] == nil {
		t.Errorf("Expected AX100 to be in the MST")
	}
}

func TestPrim_inputB_FindsAllConnections(t *testing.T) {
	// Create a new graph
	inputGraph := graph.NewGraph(strings.Split(inputB, "\n"))

	// Run Prim's algorithm
	mst := Prim(inputGraph)

	// Check the connections
	if len(mst.Vertices["AX10"].Connections) != 1 {
		t.Errorf("Expected 1 connection for AX10, got %v", len(mst.Vertices["AX10"].Connections))
	}
	if len(mst.Vertices["AX11"].Connections) != 2 {
		t.Errorf("Expected 1 connection for AX11, got %v", len(mst.Vertices["AX11"].Connections))
	}
	if len(mst.Vertices["AX12"].Connections) != 2 {
		t.Errorf("Expected 2 connections for AX12, got %v", len(mst.Vertices["AX12"].Connections))
	}
	if len(mst.Vertices["AX99"].Connections) != 2 {
		t.Errorf("Expected 1 connection for AX99, got %v", len(mst.Vertices["AX99"].Connections))
	}
	if len(mst.Vertices["AX100"].Connections) != 1 {
		t.Errorf("Expected 1 connection for AX100, got %v", len(mst.Vertices["AX100"].Connections))
	}
}
