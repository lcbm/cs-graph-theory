package main

import (
	"fmt"
	"math/rand"
)

// MaxWeight ...
const (
	minWeight = 1
	maxWeight = 30
)

var (
	a = Vertex{"A"}
	b = Vertex{"B"}
	c = Vertex{"C"}
	d = Vertex{"D"}
	e = Vertex{"E"}
	f = Vertex{"F"}
	h = Vertex{"H"}
	i = Vertex{"I"}
)

// TestA ...
func (g *Graph) TestA() {
	fmt.Printf("\n\nTesting sample A...")

	g = NewGraph()
	g.Link(&a, &a, 0)

	g.PrintVertices()
	g.PrintEdges()
	g.PrintLinks()
	g.PrintShortestPaths(&a)
}

// TestB ...
func (g *Graph) TestB() {
	fmt.Printf("\n\nTesting sample B...")

	g = NewGraph()
	g.Link(&a, &b, randomWeight())

	g.PrintVertices()
	g.PrintEdges()
	g.PrintLinks()
	g.PrintShortestPaths(&a)
}

// TestC ...
func (g *Graph) TestC() {
	fmt.Printf("\n\nTesting sample C...")

	g = NewGraph()
	g.Link(&a, &b, randomWeight())
	g.Link(&b, &c, randomWeight())
	g.Link(&c, &a, randomWeight())

	g.PrintVertices()
	g.PrintEdges()
	g.PrintLinks()
	g.PrintShortestPaths(&a)
}

// TestD ...
func (g *Graph) TestD() {
	fmt.Printf("\n\nTesting sample D...")

	g = NewGraph()
	g.Link(&a, &b, randomWeight())
	g.Link(&b, &c, randomWeight())
	g.Link(&c, &a, randomWeight())
	g.Link(&d, &a, randomWeight())
	g.Link(&d, &b, randomWeight())
	g.Link(&d, &c, randomWeight())

	g.PrintVertices()
	g.PrintEdges()
	g.PrintLinks()
	g.PrintShortestPaths(&a)
}

// TestE ...
func (g *Graph) TestE() {
	fmt.Printf("\n\nTesting sample E...")

	g = NewGraph()
	g.Link(&a, &b, randomWeight())
	g.Link(&a, &c, randomWeight())
	g.Link(&a, &d, randomWeight())
	g.Link(&a, &e, randomWeight())
	g.Link(&b, &c, randomWeight())
	g.Link(&b, &d, randomWeight())
	g.Link(&b, &e, randomWeight())
	g.Link(&c, &d, randomWeight())
	g.Link(&c, &e, randomWeight())
	g.Link(&d, &e, randomWeight())

	g.PrintVertices()
	g.PrintEdges()
	g.PrintLinks()
	g.PrintShortestPaths(&a)
}

// TestF ...
func (g *Graph) TestF() {
	fmt.Printf("\n\nTesting sample F...")

	g = NewGraph()
	g.Link(&a, &b, randomWeight())
	g.Link(&a, &c, randomWeight())
	g.Link(&a, &h, randomWeight())
	g.Link(&c, &d, randomWeight())
	g.Link(&b, &d, randomWeight())
	g.Link(&b, &e, randomWeight())
	g.Link(&b, &h, randomWeight())
	g.Link(&d, &f, randomWeight())
	g.Link(&d, &h, randomWeight())
	g.Link(&e, &h, randomWeight())
	g.Link(&e, &f, randomWeight())
	g.Link(&e, &i, randomWeight())
	g.Link(&h, &i, randomWeight())
	g.Link(&f, &i, randomWeight())

	g.PrintVertices()
	g.PrintEdges()
	g.PrintLinks()
	g.PrintShortestPaths(&a)
}

func randomWeight() int {
	return rand.Intn(maxWeight) + minWeight
}
