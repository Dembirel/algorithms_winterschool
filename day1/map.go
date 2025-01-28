package main

import "fmt"

type Graph struct {
	adj map[int][]int
	// name string
	// old int
}

func NewGraph() *Graph {
	return &Graph{adj: make(map[int][]int)}
}

func (g *Graph) AddEdge(u, v int, undirected bool) {
	/*_, ok := g.adj[u]
	if ok {
		g := NewGraph()
		g.adj[u] = append(g.adj[u], v)
		if undirected {
			g.adj[u] = append(g.adj[v], u)
		}
	} else {*/
	g.adj[u] = append(g.adj[u], v)
	if undirected {
		g.adj[v] = append(g.adj[v], u)
	}
	// }
}

func HasEdge(g *Graph, u, v int) bool {
	for _, value := range g.adj[u]{
		if value == v{
			fmt.Println(u, g.adj[u])
			fmt.Println(u, " <--> ", v)
			return true
		}
	}
	
	/*for idx, value := range g.adj {
		for _, tmp := range value {
			if v == tmp {
				fmt.Println(g[u], value)
				fmt.Println(u, " <--> ", v)
				return true
			}
		}

	}*/
	return false
}