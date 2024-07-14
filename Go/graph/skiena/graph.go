package main

type EdgeNode struct {
	y      int
	weight int
	next   *EdgeNode
}

type AdjacencyListGraph struct {
	edges          map[int]*EdgeNode // Adjacency list
	degree         map[int]int       // Degree of each vertex
	nEdges         int               // Number of edges
	directed       bool
	uniqueVertices map[int]bool
}

func NewGraph() *AdjacencyListGraph {
	return &AdjacencyListGraph{
		edges:          map[int]*EdgeNode{},
		degree:         map[int]int{},
		nEdges:         0,
		directed:       false,
		uniqueVertices: map[int]bool{},
	}
}

func (g *AdjacencyListGraph) InsertEdge(x, y int, directed bool) {
	p := &EdgeNode{}
	p.y = y
	p.next = g.edges[x]
	g.edges[x] = p
	g.degree[x]++

	g.uniqueVertices[x] = true

	if !directed {
		g.InsertEdge(y, x, true)
	} else {
		g.nEdges++
	}
}
