package main

import "fmt"

var (
	vertexA = Vertex{"A", 1}
	vertexB = Vertex{"B", 1}
	vertexC = Vertex{"C", 1}
	vertexD = Vertex{"D", 1}
	vertexE = Vertex{"E", 1}
)

// TestA ...
func (g *Graph) TestA() {
	fmt.Printf("\n\nTesting sample A...")
	g = NewGraph()
	g.AddVertex(&vertexA)
	g.Link(&vertexA, &vertexA)
	g.Show()
}

// TestB ...
func (g *Graph) TestB() {
	fmt.Printf("\n\nTesting sample B...")
	g = NewGraph()
	g.AddVertex(&vertexA)
	g.AddVertex(&vertexB)
	g.Link(&vertexA, &vertexB)
	g.Show()
}

// TestC ...
func (g *Graph) TestC() {
	fmt.Printf("\n\nTesting sample C...")
	g = NewGraph()
	g.AddVertex(&vertexA)
	g.AddVertex(&vertexB)
	g.AddVertex(&vertexC)
	g.Link(&vertexA, &vertexB)
	g.Link(&vertexB, &vertexC)
	g.Link(&vertexC, &vertexA)
	g.Show()
}

// TestD ...
func (g *Graph) TestD() {
	fmt.Printf("\n\nTesting sample D...")
	g = NewGraph()
	g.AddVertex(&vertexA)
	g.AddVertex(&vertexB)
	g.AddVertex(&vertexC)
	g.AddVertex(&vertexD)
	g.Link(&vertexA, &vertexB)
	g.Link(&vertexB, &vertexC)
	g.Link(&vertexC, &vertexA)
	g.Link(&vertexD, &vertexA)
	g.Link(&vertexD, &vertexB)
	g.Link(&vertexD, &vertexC)
	g.Show()
}

// TestE ...
func (g *Graph) TestE() {
	fmt.Printf("\n\nTesting sample E...")
	g = NewGraph()
	g.AddVertex(&vertexA)
	g.AddVertex(&vertexB)
	g.AddVertex(&vertexC)
	g.AddVertex(&vertexD)
	g.AddVertex(&vertexE)
	g.Link(&vertexA, &vertexB)
	g.Link(&vertexA, &vertexC)
	g.Link(&vertexA, &vertexD)
	g.Link(&vertexA, &vertexE)
	g.Link(&vertexB, &vertexC)
	g.Link(&vertexB, &vertexD)
	g.Link(&vertexB, &vertexE)
	g.Link(&vertexC, &vertexD)
	g.Link(&vertexC, &vertexE)
	g.Link(&vertexD, &vertexE)
	g.Show()
}
