package main

// 6
// |
// |
// 1 -----2------3------4
//        |             |
//        |             |
//         -------------5

func main() {
	graph := NewGraph()

	graph.InsertEdge(1, 2, false)
	graph.InsertEdge(2, 3, false)
	graph.InsertEdge(3, 4, false)
	graph.InsertEdge(4, 5, false)
	graph.InsertEdge(5, 2, false)
	graph.InsertEdge(1, 6, false)

	// 1. Simple BFS
	// searchData := &SearchData{
	// 	Discovered: make([]bool, len(graph.uniqueVertices)),
	// }

	// BFS(*graph, searchData, 1)

	// 2. connected components problem
	// connectedComponents(*graph)

	// 3. Two color problem
	// twocolor(*graph)

	// 4. DFS
	searchData := &SearchData{
		Discovered: make(map[int]struct{}, len(graph.uniqueVertices)),
		Processed:  make(map[int]struct{}, len(graph.uniqueVertices)),
		DFSData: DFSData{
			EntryTime: make(map[int]int, len(graph.uniqueVertices)),
			ExitTime:  make(map[int]int, len(graph.uniqueVertices)),
		},

		ProcessVertexEarly: defaultProcessVertexEarly,
		ProcessVertexLate:  defaultProcessVertexLate,
		ProcessEdge:        defaultProcessEdge,
	}
	DFS(graph, searchData, 1)
}
