package graph

import (
	"testing"
)

func TestGraphCapabilities(t *testing.T) {
	g := NewGraph(false)

	nodeA := &Node{"A"}
	nodeB := &Node{"B"}
	nodeC := &Node{"C"}
	nodeD := &Node{"D"}

	g.AddNode(nodeA)
	g.AddNode(nodeB)
	g.AddNode(nodeC)
	g.AddNode(nodeD)

	if len(g.nodes) != 4 {
		t.Fail()
	}

	g.AddEdge(nodeA, nodeC, 0)
	g.AddEdge(nodeA, nodeD, 0)
	g.AddEdge(nodeB, nodeD, 0)

	if len(g.edges[*nodeA]) != 2 {
		t.Fail()
	}

	if len(g.edges[*nodeB]) != 1 {
		t.Fail()
	}

	if len(g.edges[*nodeC]) != 1 {
		t.Fail()
	}

	if len(g.edges[*nodeD]) != 2 {
		t.Fail()
	}

	var s = `A -> (C, 0.00) (D, 0.00)
B -> (D, 0.00)
C -> (A, 0.00)
D -> (A, 0.00) (B, 0.00)
`

	if g.String() != s {
		t.Fail()
	}
}

func TestNodeMethods(t *testing.T) {
	node := &Node{
		Data: 42,
	}

	if node.String() != "42" {
		t.Fail()
	}
}

func TestSortEdges(t *testing.T) {
	var edges = edges{
		edge{from: &Node{"A"}, to: &Node{"B"}, weight: 4},
		edge{from: &Node{"A"}, to: &Node{"C"}, weight: 1},
		edge{from: &Node{"A"}, to: &Node{"D"}, weight: 2},
	}

	edges.Sort()

	if edges[0].weight != 1 {
		t.Fail()
	}

	if edges[1].weight != 2 {
		t.Fail()
	}

	if edges[2].weight != 4 {
		t.Fail()
	}
}
