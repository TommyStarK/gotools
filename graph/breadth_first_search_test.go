package graph

import (
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
	g := NewGraph(true)

	nodeA := &Node{"A"}
	nodeB := &Node{"B"}
	nodeC := &Node{"C"}
	nodeD := &Node{"D"}
	nodeE := &Node{"E"}
	nodeF := &Node{"F"}
	nodeG := &Node{"G"}
	nodeH := &Node{"H"}

	g.AddEdge(nodeA, nodeB, 3)
	g.AddEdge(nodeA, nodeC, 6)
	g.AddEdge(nodeB, nodeC, 4)
	g.AddEdge(nodeB, nodeD, 4)
	g.AddEdge(nodeB, nodeE, 11)
	g.AddEdge(nodeC, nodeD, 8)
	g.AddEdge(nodeC, nodeG, 11)
	g.AddEdge(nodeD, nodeE, -4)
	g.AddEdge(nodeD, nodeF, 5)
	g.AddEdge(nodeD, nodeG, 2)
	g.AddEdge(nodeE, nodeH, 9)
	g.AddEdge(nodeF, nodeH, 1)
	g.AddEdge(nodeG, nodeH, 2)

	witness := make(map[string]bool, len(g.nodes))
	for _, value := range []string{"A", "B", "C", "D", "E", "F", "G", "H"} {
		witness[value] = false
	}

	g.BreadthFirstSearch(nodeA, func(node *Node) {
		witness[node.String()] = true
	})

	for _, visited := range witness {
		if !visited {
			t.Fail()
		}
	}
}
