package graph

import (
	"strconv"
	"strings"
)

type Graph struct {
	Vertices map[string]*Vertex
	Count    int
}

// EmptyGraph creates a new empty graph
func EmptyGraph() *Graph {
	return &Graph{
		Vertices: make(map[string]*Vertex),
	}
}

// NewGraph creates a new graph based on the adjacencyList provided.
// The first element is expected to be a comma separated string of vertex values
// The rest of the elements describe the connections to be added to each vertex.
//
// Big-O: O(n^2) because it loops through each vertex and each connection of the vertex
func NewGraph(adjacencyList []string) *Graph {
	if len(adjacencyList) == 0 {
		panic("Adjacency list is empty")
	}

	// Create a new graph
	graph := &Graph{
		Vertices: make(map[string]*Vertex),
	}

	// Parse the Vertices
	for _, key := range strings.Split(adjacencyList[0], ",") {
		graph.AddVertex(key)
	}

	// Add connections
	// Start at i = 1 bc the first element is not a connection
	for i := 1; i < len(adjacencyList); i++ {
		// Parse the vertex value and the connections
		vertexAndConnections := strings.Split(adjacencyList[i], ",")
		vertex := graph.GetVertex(vertexAndConnections[0])

		// Add each of the connections
		for j := 1; j < len(vertexAndConnections); j++ {
			// Parse the connection
			connection := strings.Split(vertexAndConnections[j], ":")
			if len(connection) != 2 {
				panic("Invalid connection")
			}

			// Parse weight
			weight, err := strconv.Atoi(connection[1])
			if err != nil {
				panic(err)
			}

			// Add the connecting vertex if it doesn't exist, and add the connection
			if graph.GetVertex(connection[0]) == nil {
				panic("Connection vertex doesn't exist")
			}
			graph.AddConnection(vertex.Key, connection[0], weight)
		}
	}

	return graph
}

// AddVertex creates a new Vertex with the provided key and adds it to the graph
// If the key already exists, nothing is done
//
// Big-O: O(1) because it just creates a new vertex
func (g *Graph) AddVertex(key string) *Vertex {
	// Check for duplicate
	if _, ok := g.Vertices[key]; ok {
		return g.Vertices[key]
	}

	// Create vertex
	vertex := NewVertex(key)
	g.Vertices[key] = vertex
	g.Count++
	return vertex
}

// GetVertex returns the vertex with the provided key, or nil if it doesn't exist
//
// Big-O: O(1) because it just returns the vertex
func (g *Graph) GetVertex(key string) *Vertex {
	return g.Vertices[key]
}

// GetKeys returns a slice of all the keys in the graph, in no particular order
//
// Big-O: O(n) because it loops through each vertex
func (g *Graph) GetKeys() []string {
	keys := make([]string, 0, len(g.Vertices))
	for key := range g.Vertices {
		keys = append(keys, key)
	}
	return keys
}

// AddConnection adds a connection between the two vertices with the provided keys
// If either of the vertices doesn't exist, nothing is done
//
// Big-O: O(1) because it just adds a connection to the vertex
func (g *Graph) AddConnection(key1, key2 string, weight int) {
	// Check if vertices exist
	vertex1, ok := g.Vertices[key1]
	if !ok {
		return
	}
	vertex2, ok := g.Vertices[key2]
	if !ok {
		return
	}

	// Add the connection
	vertex1.AddConnection(key2, weight)
	vertex2.AddConnection(key1, weight)
}
