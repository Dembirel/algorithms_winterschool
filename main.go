package main

import "fmt"

func main() {
	var a = Graph{make(map[int][]int)}
	var b = Graph{make(map[int][]int)}
	var c = Graph{make(map[int][]int)}
	var d = Graph{make(map[int][]int)}
	a.AddEdge(1, 2, false)
	b.AddEdge(2, )
	c.AddEdge()
	d.AddEdge()
	if HasEdge(&a, 1, 2) {

	}
	fmt.Println("Hello world!")
}
