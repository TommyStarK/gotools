package graph

import (
	"fmt"
	"sort"
	"sync"
)

type edge struct {
	from   *Node
	to     *Node
	weight float64
}

type edges []edge

func (e edges) Sort() {
	sort.Slice(e, func(a, b int) bool {
		return e[a].weight < e[b].weight
	})
}

type Node struct {
	Data interface{}
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Data)
}

type Graph struct {
	digraph bool
	edges   map[Node]edges
	nodes   []*Node

	mutex sync.RWMutex
}

func NewGraph(directed bool) *Graph {
	return &Graph{
		digraph: directed,
		edges:   make(map[Node]edges),
		nodes:   make([]*Node, 0),
		mutex:   sync.RWMutex{},
	}
}

func (g *Graph) AddEdge(from, to *Node, weight float64) {
	g.mutex.Lock()

	g.edges[*from] = append(g.edges[*from], edge{from: from, to: to, weight: weight})
	if !g.digraph {
		g.edges[*to] = append(g.edges[*to], edge{from: to, to: from, weight: weight})
	}

	g.mutex.Unlock()
}

func (g *Graph) AddNode(node *Node) {
	g.mutex.Lock()
	g.nodes = append(g.nodes, node)
	g.mutex.Unlock()
}

func (g *Graph) String() (s string) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].String()
		edges := g.edges[*g.nodes[i]]

		if len(edges) > 0 {
			s += " ->"
		}

		for j := 0; j < len(edges); j++ {
			s += " (" + edges[j].to.String() + ", " + fmt.Sprintf("%.2f", edges[j].weight) + ")"
		}

		s += "\n"
	}

	return
}
