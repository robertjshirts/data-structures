package mst

import "github.com/robertjshirts/data-structures/graph"

// Prim is an implementation of Prim's algorithm for finding the minimum spanning tree of a graph.
//
// Big-O: O(n^2) because it loops through each vertex in the graph and each connection of the vertex
func Prim(inputGraph *graph.Graph) *graph.Graph {
	// Validate graph
	if inputGraph == nil || len(inputGraph.Vertices) == 0 {
		return graph.EmptyGraph()
	}

	// Create a new graph
	mst := graph.EmptyGraph()

	// Add the first vertex to the MST
	firstVertex := inputGraph.GetVertex(inputGraph.GetKeys()[0])
	mst.AddVertex(firstVertex.Key)

	// Loop until all vertices are in the MST
	for len(mst.Vertices) < len(inputGraph.Vertices) {
		// Find the closest vertex to the MST that is not already in the MST
		from, to, weight := findClosestConnection(inputGraph, mst)

		// Add the closest vertex to the MST
		mst.AddVertex(to)

		// Add the connection to the MST
		mst.AddConnection(from, to, weight)
	}

	return mst
}

// findClosestUniqueVertexKey finds the closest vertex to the MST that is not already in the MST, returns from, to, and weight of the connection
//
// Big-O: O(n^2) because it loops through each vertex in the MST and each connection of the vertex
func findClosestConnection(inputGraph *graph.Graph, mst *graph.Graph) (string, string, int) {
	from := ""
	to := ""
	currentWeight := 999999
	// Look through each vertex in the MST
	for key := range mst.Vertices {
		// Look through each connection of the vertex
		for connectionKey, weight := range inputGraph.GetVertex(key).Connections {
			// If the connection is not in the MST, and it is the closest one found so far, save it
			if _, ok := mst.Vertices[connectionKey]; !ok && weight < currentWeight {
				from = key
				to = connectionKey
				currentWeight = weight
			}
		}
	}
	return from, to, currentWeight
}

// Weight returns the total weight of the graph
//
// Big-O: O(n^2) because it loops through each vertex in the graph and each connection of the vertex
func Weight(g *graph.Graph) int {
	weight := 0
	for _, vertex := range g.Vertices {
		for _, w := range vertex.Connections {
			weight += w
		}
	}
	// The weight is counted twice, so divide by 2
	return weight / 2
}
