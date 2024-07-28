package main

// Generally, 3 cases (checked by process vertex late	):
// - root cute node with 2 or more children
// - bridge cut node - if the earliest reachable vertex from the vertex is the vertex itself
// then, deleting the its parent will disconnect the graph. The v itself is also an articulation vertex
// unless it is a leaf.
// - parent cut node - if the earliest reachable vertex from the vertex is the parent of the vertex,
// then deleting the parent will disconnect the graph. The parent is an articulation vertex.
func articulationVertices(g *AdjacencyListGraph) {

	// // earlieast reachable vertex from each vertex
	// reacheableAncestor := make(map[int]int, len(g.uniqueVertices))
	// // DFS tree out degree of each vertex
	// treeOutDegree := make(map[int]int, len(g.uniqueVertices))

	// searchData := &SearchData{}

	// processVertexEarly := func(v int) {
	// 	reacheableAncestor[v] = v
	// }

	// processEdge := func(x, y int) {
	// 	class := edgeClassification(x, y)

	// 	if class == TreeEdge {
	// 		treeOutDegree[x]++
	// 	}

	// 	if class == BackEdge && searchData.Parent[x] != y {
	// 		if searchData.EntryTime[y] < searchData.EntryTime[reacheableAncestor[x]] {
	// 			reacheableAncestor[x] = y
	// 		}
	// 	}
	// }

	// processVertexLate := func(v int) {
	// 	var (
	// 		isRoot     false // flag indicating if the vertex is the root of the DFS tree
	// 		timeV      int   // earliest reachable time of the vertex
	// 		timeParent int   // earliest reachable time of the parent of the vertex
	// 	)

	// 	if searchData.Parent[v] == 0 { // test if v is the root of the DFS tree
	// 		if treeOutDegree[v] > 1 {
	// 			println("Root articulation vertex:", v)
	// 		}
	// 		return
	// 	}

	// 	// is parent of v the root of the DFS tree?
	// 	isParentRoot := searchData.Parent[searchData.Parent[v]] == 0
	// 	if reacheableAncestor[v] == searchData.Parent[v] && !isParentRoot {
	// 		println("Parent articulation vertex:", searchData.Parent[v])
	// 	}

	// 	if reacheableAncestor[v] == v {
	// 		println("Bridge articulation vertex:", searchData.Parent[v])

	// 		if treeOutDegree[v] > 0 { // test uf v is not a leaf
	// 			println("Bridge articulation vertex:", v)
	// 		}
	// 	}

	// 	timeV = searchData.EntryTime[reacheableAncestor[v]]
	// 	timeParent = searchData.EntryTime[reacheableAncestor[searchData.Parent[v]]]

	// 	if timeV < timeParent {
	// 		reacheableAncestor[searchData.Parent[v]] = reacheableAncestor[v]
	// 	}
	// }
}

func edgeClassification(x, y int) Class {
	// TODO:
	return TreeEdge
}
