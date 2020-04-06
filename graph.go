package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	// ErrVertexExists is returned when adding a vertex with a label that already exists.
	ErrVertexExists = errors.New("vertex already exists, try a different label")

	// ErrEdgeExists is returned when adding a vertex with a label that already exists.
	ErrEdgeExists = errors.New("edge already exists, try a different label")
)

// Graph ...
type Graph struct {
	Vertices map[*Vertex]struct{}
	Edges    map[*Edge]struct{}
}

// Vertex ...
type Vertex struct {
	Label string
}

// Edge ...
type Edge struct {
	From   *Vertex
	To     *Vertex
	Weight int
}

// NewGraph ...
func NewGraph() *Graph {
	return &Graph{
		Vertices: make(map[*Vertex]struct{}, 0),
		Edges:    make(map[*Edge]struct{}, 0),
	}
}

// Link ...
func (g *Graph) Link(from, to *Vertex, weight int) error {
	// Add source vertex to vertices slice.
	err := g.AddVertex(from)
	if err != nil {
		fmt.Errorf("link between %v and %v failed: %w", from, to, err)
	}
	// Add destination vertex to vertices slice.
	err = g.AddVertex(to)
	if err != nil {
		fmt.Errorf("link between %v and %v failed: %w", from, to, err)
	}
	// Add edge from source to destination.
	err = g.AddEdge(from, to, weight)
	if err != nil {
		fmt.Errorf("link between %v and %v failed: %w", from, to, err)
	}
	// Add edge from destination to source (undirected graph).
	err = g.AddEdge(to, from, weight)
	if err != nil {
		fmt.Errorf("link between %v and %v failed: %w", to, from, err)
	}
	return nil
}

// AddVertex ...
func (g *Graph) AddVertex(v *Vertex) error {
	// Check if vertex exists before proceeding.
	if _, ok := g.Vertices[v]; ok {
		return fmt.Errorf("failed to add vertex: %s", ErrVertexExists)
	}
	// Add vertex to vertices slice.
	g.Vertices[v] = struct{}{}
	return nil
}

// AddEdge ...
func (g *Graph) AddEdge(from, to *Vertex, weight int) error {
	// Create new edge
	edge := &Edge{
		From:   from,
		To:     to,
		Weight: weight,
	}
	// Check if edge exists before proceeding
	if _, ok := g.Edges[edge]; ok {
		return fmt.Errorf("failed to add edge: %s", ErrEdgeExists)
	}
	// Add edge to root slice
	g.Edges[edge] = struct{}{}
	return nil
}

// PrintVertices prints all vertices from the graph.
func (g *Graph) PrintVertices() {
	vertices := fmt.Sprintf("\nVertices:\n\t")
	for v := range g.Vertices {
		vertices += fmt.Sprintf("%s, ", v.Label)
	}
	vertices = strings.TrimRight(vertices, ", ")
	fmt.Println(vertices)
}

// PrintEdges prints all edges from the graph.
func (g *Graph) PrintEdges() {
	edges := fmt.Sprintf("\nEdges:")
	for e := range g.Edges {
		edges += fmt.Sprintf("\n\t%s -> %s (%s)", e.From.Label, e.To.Label, strconv.Itoa(e.Weight))
	}
	fmt.Println(edges)
}

// PrintLinks prints all links from the graph.
func (g *Graph) PrintLinks() {
	links := fmt.Sprintf("\nLinks:")
	for v := range g.Vertices {
		links += fmt.Sprintf("\n\t%s: ", v.Label)
		for e := range g.Edges {
			if e.From.Label == v.Label {
				links += fmt.Sprintf("%s -> ", e.To.Label)
			}
		}
		links = strings.TrimRight(links, " -> ")
	}
	fmt.Println(links)
}
