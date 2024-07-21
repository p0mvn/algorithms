package main

import "fmt"

func DFS(graph *AdjacencyListGraph, searchData *SearchData, startV int) error {

	// TODO: check for termination

	v := startV

	vEdges, ok := graph.edges[v]
	if !ok {
		return fmt.Errorf("Vertex %d not found in graph", v)
	}

	searchData.Discovered[v] = struct{}{}
	searchData.Time++
	searchData.EntryTime[v] = searchData.Time

	searchData.ProcessVertexEarly(v)

	for p := vEdges; p != nil; p = p.next {
		y := p.y

		_, ok := searchData.Discovered[y]
		if !ok {
			searchData.ProcessEdge(v, y)
			DFS(graph, searchData, y)
		} else if _, ok := searchData.Processed[y]; !ok || graph.directed {
			searchData.ProcessEdge(v, y)

		}

		// TODO: check for termination
	}

	searchData.ProcessVertexLate(v)

	searchData.Time++
	searchData.ExitTime[v] = searchData.Time

	searchData.Processed[v] = struct{}{}

	return nil
}
