package main

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
