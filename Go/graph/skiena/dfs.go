package main

import "fmt"

func DFS(graph *AdjacencyListGraph, searchData *SearchData, startV int) error {
	if searchData.Finished {
		fmt.Println("Finished search")
		return nil
	}

	v := startV

	vEdges, ok := graph.edges[v]
	if !ok {
		return fmt.Errorf("Vertex %d not found in graph", v)
	}

	searchData.Discovered[v] = struct{}{}
	searchData.Time++
	searchData.EntryTime[v] = searchData.Time

	searchData.ProcessVertexEarly(v)

	p := vEdges
	for p != nil {
		y := p.y

		_, ok := searchData.Discovered[y]
		if !ok {
			searchData.Parent[y] = v
			searchData.ProcessEdge(v, y)
			DFS(graph, searchData, y)
		} else if _, ok := searchData.Processed[y]; !ok || graph.directed {
			searchData.ProcessEdge(v, y)
		}

		if searchData.Finished {
			fmt.Println("Finished search")
			return nil
		}

		p = p.next
	}

	searchData.ProcessVertexLate(v)

	searchData.Time++
	searchData.ExitTime[v] = searchData.Time

	searchData.Processed[v] = struct{}{}

	return nil
}
