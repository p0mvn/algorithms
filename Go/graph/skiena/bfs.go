package main

func BFS(graph AdjacencyListGraph, searchData *SearchData, startV int) {
	queue := make([]int, 0)

	queue = append(queue, startV)

	processed := make(map[int]struct{}, len(graph.uniqueVertices))

	parent := make(map[int]int, len(graph.uniqueVertices))

	searchData.Discovered[startV] = struct{}{}
	processed[startV] = struct{}{}

	for len(queue) > 0 {
		curV := queue[0]
		queue = queue[1:]

		// Process the current vertex
		searchData.ProcessVertexEarly(curV)

		for p := graph.edges[curV]; p != nil; p = p.next {
			y := p.y

			_, ok := processed[y]
			if !ok || graph.directed {
				searchData.ProcessEdge(curV, y)
			}

			_, ok = searchData.Discovered[y]
			if !ok {
				queue = append(queue, y)
				searchData.Discovered[y] = struct{}{}
				parent[y] = curV
			}
		}

		searchData.ProcessVertexLate(curV)
	}

	printParents(parent)
}

func defaultProcessVertexEarly(v int) {
	println("Processed vertex early:", v)
}

func defaultProcessVertexLate(v int) {
	println("Processed vertex late:", v)
}

func defaultProcessEdge(x, y int) {
	println("Processed edge:", x, y)
}

func printParents(parents map[int]int) {
	println("Parents:")
	for i, p := range parents {
		println(i, ":", p)
	}
}
