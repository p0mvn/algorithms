package main

func BFS(graph AdjacencyListGraph, searchData *SearchData, startV int) {
	queue := make([]int, 0)

	queue = append(queue, startV)

	processed := make([]bool, len(graph.uniqueVertices))

	parent := make([]int, len(graph.uniqueVertices))

	searchData.Discovered[startV-1] = struct{}{}
	processed[startV-1] = true

	for len(queue) > 0 {
		curV := queue[0]
		queue = queue[1:]

		// Process the current vertex
		searchData.ProcessVertexEarly(curV)

		for p := graph.edges[curV]; p != nil; p = p.next {
			y := p.y
			yIndex := y - 1

			if !processed[yIndex] || graph.directed {
				searchData.ProcessEdge(curV, y)
			}

			_, ok := searchData.Discovered[yIndex]
			if !ok {
				queue = append(queue, y)
				searchData.Discovered[yIndex] = struct{}{}
				parent[yIndex] = curV
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

func printParents(parents []int) {
	println("Parents:")
	for i, p := range parents {
		println(i+1, ":", p)
	}
}
