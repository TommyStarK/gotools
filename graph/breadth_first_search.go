package graph

func (g *Graph) BreadthFirstSearch(from *Node, f func(*Node)) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	q := make([]*Node, 0, len(g.nodes))
	visited := make(map[Node]bool, len(g.nodes))

	q = append(q, from)
	for {
		if len(q) == 0 {
			break
		}

		node := q[0]
		q = q[1:]

		if visited[*node] {
			continue
		}

		visited[*node] = true
		neighbours := g.edges[*node]

		for i := 0; i < len(neighbours); i++ {
			nextNode := neighbours[i].to
			if !visited[*nextNode] {
				q = append(q, nextNode)
			}
		}

		if f != nil {
			f(node)
		}
	}
}
