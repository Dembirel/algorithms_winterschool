package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
	var a = Graph{make(map[int][]int)}
	a.AddEdge(1, 2, true)
	a.AddEdge(1, 3, true)
	a.AddEdge(2, 4, true)
	a.AddEdge(2, 5, true)
	a.AddEdge(3, 6, true)
	a.AddEdge(5, 6, true)
	HasEdge(&a, 3, 1)
	HasEdge(&a, 2, 4)
	HasEdge(&a, 5, 2)	

}
