package graph

type Vertex struct {
	Key         string
	Connections map[string]int
}

// NewVertex creates a new vertex with the given key.
//
// Big-O: O(1) because it creates a new vertex
func NewVertex(key string) *Vertex {
	return &Vertex{
		Key:         key,
		Connections: make(map[string]int),
	}
}

// AddConnection adds a connection to the vertex with the given key and weight.
//
// Big-O: O(1) because it adds a connection to the map
func (v *Vertex) AddConnection(key string, weight int) {
	v.Connections[key] = weight
}
