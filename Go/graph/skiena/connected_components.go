package main

func connectedComponents(graph AdjacencyListGraph) {

	searchData := &SearchData{
		Discovered: make(map[int]struct{}, len(graph.uniqueVertices)),
	}

	c := 0

	for i := range graph.uniqueVertices {
		_, ok := searchData.Discovered[i-1]
		if !ok {
			c++
			println("Component", c)
			BFS(graph, searchData, i+1)
		}
	}

}
