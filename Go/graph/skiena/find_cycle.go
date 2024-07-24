package main

import "fmt"

func findCycle(graph *AdjacencyListGraph) bool {

	searchData := &SearchData{
		Discovered: make(map[int]struct{}, len(graph.uniqueVertices)),
		Processed:  make(map[int]struct{}, len(graph.uniqueVertices)),
		DFSData: DFSData{
			EntryTime: make(map[int]int, len(graph.uniqueVertices)),
			ExitTime:  make(map[int]int, len(graph.uniqueVertices)),
		},
		Parent: make(map[int]int, len(graph.uniqueVertices)),

		ProcessVertexEarly: defaultProcessVertexEarly,
		ProcessVertexLate:  defaultProcessVertexLate,
		ProcessEdge:        defaultProcessEdge,
	}

	findCycleEdge := func(x, y int) {
		if searchData.Parent[y] != x && (graph.directed || searchData.Parent[x] != y) {
			fmt.Println("Found cycle:", x, y)
			searchData.Finished = true
		} else {
			defaultProcessEdge(x, y)
		}
	}

	searchData.ProcessEdge = findCycleEdge

	for vertex := range graph.uniqueVertices {
		if _, ok := searchData.Discovered[vertex]; !ok {
			if err := DFS(graph, searchData, vertex); err != nil {
				return false
			}
		}
	}

	return false
}
