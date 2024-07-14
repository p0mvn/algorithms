package main

func main() {
	graph := NewGraph()

	graph.InsertEdge(1, 2, false)
	graph.InsertEdge(2, 3, false)
	graph.InsertEdge(3, 4, false)
	graph.InsertEdge(4, 5, false)
	graph.InsertEdge(5, 2, false)
	graph.InsertEdge(1, 6, false)

	BFS(*graph, 1)
}
