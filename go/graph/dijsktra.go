// This implementation is heavily inspired on @vaidehijoshi article on Dijkstra:
// - medium.com/basecs/finding-the-shortest-path-with-a-little-help-from-dijkstra-613149fbdc8e

package main

import (
	"fmt"
	"sort"
)

// infinite ...
const infinite = int(^uint(0) >> 1)

// DijkstraTable ...
type DijkstraTable struct {
	Vertex *Vertex
	Weight int
}

// NewTable ...
func (g *Graph) NewTable() map[*Vertex]int {
	return make(map[*Vertex]int, 0)
}

// InitializeTable ...
func (g *Graph) InitializeTable(table map[*Vertex]int, origin *Vertex) map[*Vertex]int {
	// Initialize all shortest paths as inifinity.
	for v := range g.Vertices {
		table[v] = infinite
	}
	// Assign 0 to origin vertex.
	table[origin] = 0
	return table
}

// Dijkstra ...
func (g *Graph) Dijkstra(origin *Vertex) map[*Vertex]int {
	// Create a table to keep track of the shortest known distance to
	// every vertex in the graph, as well as the previous vertex that
	// we came from -- before checking the current vertex.
	table := g.NewTable()
	// Initialize table values with infinity to start out the shortest paths
	// (considering we don't even know if all vertices are reachable), except
	// the path from the origin vertex to iteself (which is initialized as 0).
	table = g.InitializeTable(table, origin)
	// Create a slice to keep track of which vertices were already visited.
	visited := make(map[*Vertex]struct{}, 0)
	for range g.Vertices {
		// Get closest non visited vertex.
		vertex := GetClosestVertex(table, visited)
		// Get vertex's neighboring edges.
		edges := g.GetNeighborEdges(vertex)
		for _, e := range edges {
			// Calculate distance from the current vertex to its neighbors.
			distance := table[vertex] + e.Weight
			// Check if the calculated distance is less than the currently-known
			// shortest distance for current neighboring vertex.
			if distance < table[e.To] {
				// Update table with the new “shortest distance”.
				table[e.To] = distance
			}
		}
		// Mark current vertex as visited.
		visited[vertex] = struct{}{}
		// visited = append(visited, vertex)
	}
	return table
}

// GetClosestVertex returns the closest unvisited vertex from the table.
func GetClosestVertex(table map[*Vertex]int, visited map[*Vertex]struct{}) *Vertex {
	var unvisited []*DijkstraTable
	// Verify if the vertex has been visited already.
	for vertex, weight := range table {
		if _, ok := visited[vertex]; !ok {
			unvisited = append(unvisited, &DijkstraTable{vertex, weight})
		}
	}
	// Sort slice to retrieve the vertex with the loweest weight from the table.
	sort.Slice(unvisited, func(i, j int) bool {
		return unvisited[i].Weight < unvisited[j].Weight
	})
	return unvisited[0].Vertex
}

// GetNeighborEdges returns all the Edges connecting to the vertex's neighbors.
func (g *Graph) GetNeighborEdges(vertex *Vertex) []*Edge {
	edges := make([]*Edge, 0)
	for e := range g.Edges {
		if e.From == vertex {
			edges = append(edges, e)
		}
	}
	return edges
}

// PrintShortestPaths prints the shortest paths from origin vertex to all others.
func (g *Graph) PrintShortestPaths(origin *Vertex) {
	spaths := fmt.Sprintf("\nShortest Paths:")
	for v, w := range g.Dijkstra(origin) {
		spaths += fmt.Sprintf("\n\tfrom %s to %s: %d", origin.Label, v.Label, w)
	}
	fmt.Println(spaths)
}
