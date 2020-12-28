package main

import "fmt"

type Graph map[int][]int

func NewGraph() Graph {
	return Graph{}
}

func (g Graph) AddVertex(v int) {
	g[v] = []int{}
}

func (g Graph) AddEdge(v1, v2 int) {
	g[v1] = append(g[v1], v2)
	g[v2] = append(g[v2], v1)
}

func (g Graph) Graph() map[int][]int {
	return g
}

func (g Graph) Entry(v int) []int {
	return g[v]
}

func main() {
	var n, m, c, d int
	fmt.Scan(&n, &m)
	a := make([]int, n)
	b := make([]int, n)
	diff := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&b[i])
		diff[i] = b[i] - a[i]
	}

	graph := NewGraph()
	for i := 0; i < n; i++ {
		graph.AddVertex(i)
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&c, &d)
		graph.AddEdge(c-1, d-1)
	}
	fmt.Println(graph.Graph())
}
