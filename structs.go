package main

import "fmt"

type Stack struct {
	data []int
}

func newStack(s *Stack) *Stack {
	return &Stack{data: make([]int, 0)}
}

func (s *Stack) Push(a int) {
	s.data = append(s.data, a)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	topIndex := len(s.data) - 1
	val := s.data[topIndex]
	s.data = s.data[:topIndex]
	return val, true
}

func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}

type simpleQueue []int

func (q *simpleQueue) EnQueue(x int) {
	*q = append(*q, x)
}

func (q *simpleQueue) DeQueue() (int, bool) {
	if len(*q) == 0 {
		return 0, false
	}
	val := (*q)[0]
	*q = (*q)[1:]
	return val, true
}

type Graph struct {
	adj map[int][]int
	// name string
	// old int
}

func NewGraph() *Graph {
	return &Graph{adj: make(map[int][]int)}
}

func (g *Graph) AddEdge(u, v int, undirected bool) {
	val, ok := g.adj[u]
	g.adj[u] = append(g.adj[u], v)
	if undirected {
		g.adj[u] = append(g.adj[v], u)
	}

	if ok {
		g := NewGraph()
		g.adj[u] = append(g.adj[u], v)
		if undirected {
			g.adj[u] = append(g.adj[v], u)
		}
	}

}

func HasEdge(g *Graph, u, v int) bool {
	// for i = 0, i < len(g.adj); i++ {

	// }
	for idx, value := range g.adj {
		fmt.Println(idx, value)
		for _, tmp := range value {
			if v == tmp {
				return true
			}
		}

	}
	return false
}
