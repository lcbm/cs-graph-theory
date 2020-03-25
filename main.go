package main

import (
	"errors"
	"fmt"
)

var (
	// ErrVertexExists is returned when adding a vertex with a label that already exists
	ErrVertexExists = errors.New("vertex already exists, try a different label")

	// ErrEdgeExists is returned when adding a vertex with a label that already exists
	ErrEdgeExists = errors.New("edge already exists, try a different label")
)

// Vertex ...
type Vertex struct {
	Label  string
	Weigth float64
}

// Graph ...
type Graph struct {
	Vertices map[*Vertex]struct{}
	Edges    map[*Vertex]map[*Vertex]struct{}
}

// NewGraph ...
func NewGraph() *Graph {
	return &Graph{
		Vertices: make(map[*Vertex]struct{}, 0),
		Edges:    make(map[*Vertex]map[*Vertex]struct{}, 0),
	}
}

// Link ...
func (g *Graph) Link(src, dst *Vertex) {
	// Add src vertex to vertices slice
	err := g.AddVertex(src)
	if err != g.AddVertex(src) {
		fmt.Errorf("failed link vertex %v to %v: %w", src, dst, err)
	}
	// Add dst vertex to vertices slice
	err = g.AddVertex(dst)
	if err != nil {
		fmt.Errorf("failed link vertex %v to %v: %w", src, dst, err)
	}
	// Add and link src to dst
	err = g.AddEdge(src, dst)
	if err != nil {
		fmt.Errorf("failed link vertex %v to %v: %w", src, dst, err)
	}
	// Add and link dst to src (undirected graph)
	err = g.AddEdge(dst, src)
	if err != nil {
		fmt.Errorf("failed link vertex %v to %v: %w", dst, src, err)
	}
}

// AddVertex ...
func (g *Graph) AddVertex(v *Vertex) error {
	// Check if vertex exists before proceeding
	if _, exists := g.Vertices[v]; exists {
		return fmt.Errorf("failed to add vertex: %s", ErrVertexExists)
	}
	// Add vertex to vertices slice
	g.Vertices[v] = struct{}{}
	return nil
}

// AddEdge ...
func (g *Graph) AddEdge(src, dst *Vertex) error {
	// Check if edge exists before proceeding
	if _, exists := g.Edges[src][dst]; exists {
		return fmt.Errorf("failed to add edge: %s", ErrEdgeExists)
	}
	// Check if vertex is in root slice
	if _, exists := g.Edges[src]; !exists {
		g.Edges[src] = make(map[*Vertex]struct{}, 0)
	}
	// Link src to dst
	g.Edges[src][dst] = struct{}{}
	return nil
}

// Show ...
func (g *Graph) Show() {
	for vertex, links := range g.Edges {
		fmt.Printf("\nLink: %s -> ", vertex.Label)
		for v := range links {
			fmt.Printf("%s ", v.Label)
		}
	}
}

func main() {
	g := NewGraph()
	g.TestA()
	g.TestB()
	g.TestC()
	g.TestD()
	g.TestE()
}
