package graph

func (g *Graph) DepthFirstSearch(from *Node, f func(*Node)) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	stack := make([]*Node, 0, len(g.nodes))
	visited := make(map[Node]bool, len(g.nodes))

	stack = append(stack, from)
	for {
		if len(stack) == 0 {
			break
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[*node] {
			continue
		}

		visited[*node] = true
		neighbours := g.edges[*node]

		for i := 0; i < len(neighbours); i++ {
			nextNode := neighbours[i].to
			stack = append(stack, nextNode)
		}

		if f != nil {
			f(node)
		}
	}
}
